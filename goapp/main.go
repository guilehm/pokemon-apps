package main

import (
	"context"
	"encoding/json"
	"fmt"
	"goapp/db"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
)

const POKEMON_API_BASE_URL = "https://pokeapi.co/api/v2"
const POKEMON_API_LIST_URL = "https://pokeapi.co/api/v2/pokemon/"

type Pokemon struct {
	Id        int32     `bson:"id" json:"id"`
	Name      string    `bson:"name" json:"name"`
	Height    int32     `bson:"height,omitempty" json:"height,omitempty"`
	Weight    int32     `bson:"weight,omitempty" json:"weight,omitempty"`
	DateAdded time.Time `bson:"date_added,omitempty" json:"date_added,omitempty"`
	Sprites   struct {
		Other struct {
			OfficialArtwork struct {
				FrontDefault string `bson:"front_default" json:"front_default"`
			} `bson:"official-artwork" json:"official-artwork"`
		} `bson:"other" json:"other"`
	} `bson:"sprites" json:"sprites"`
}

type PokemonApiListResult struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
type PokemonApiListResponse struct {
	Count    int                    `json:"count"`
	Next     string                 `json:"next,omitempty"`
	Previous string                 `json:"previous,omitempty"`
	Results  []PokemonApiListResult `json:"results"`
}

func getPokemonDetail(id string) Pokemon {
	fmt.Printf("Fetching %q\n", id)
	endpoint := POKEMON_API_LIST_URL + id
	resp, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}
	var pokemon Pokemon
	err = json.NewDecoder(resp.Body).Decode(&pokemon)
	if err != nil {
		panic(err)
	}
	return pokemon
}

func handleApiErrors(w http.ResponseWriter, status int, message string) {
	if message == "" {
		message = http.StatusText(status)
	}

	jsonResponse, _ := json.Marshal(struct {
		Error string `json:"error"`
	}{message})
	w.WriteHeader(status)
	w.Write(jsonResponse)
}

func pokemonApiDetail(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	id, ok := vars["id"]

	if !ok {
		handleApiErrors(w, http.StatusBadRequest, "")
		return
	}

	_, idError := strconv.Atoi(id)
	if idError != nil {
		handleApiErrors(w, http.StatusBadRequest, "")
		return
	}

	endpoint := POKEMON_API_LIST_URL + id
	resp, err := http.Get(endpoint)
	if err != nil {
		handleApiErrors(w, resp.StatusCode, "")
		return
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	pokemon := Pokemon{}
	err = json.NewDecoder(resp.Body).Decode(&pokemon)
	jsonResponse, _ := json.Marshal(pokemon)

	if err != nil {
		handleApiErrors(w, http.StatusInternalServerError, "")
		return
	}

	w.Write(jsonResponse)
}

func pokemonApiList(w http.ResponseWriter, req *http.Request) {
	const defaultLimit = "20"
	const defaultOffset = "0"
	w.Header().Set("Content-Type", "application/json")
	limit := req.FormValue("limit")
	offset := req.FormValue("offset")

	if limit == "" {
		limit = defaultLimit
	}

	if offset == "" {
		offset = defaultOffset
	}

	_, limitIntErr := strconv.Atoi(limit)
	_, offsetIntErr := strconv.Atoi(offset)

	if limitIntErr != nil || offsetIntErr != nil {
		handleApiErrors(w, http.StatusBadRequest, "")
		return
	}

	endpoint, err := url.Parse(POKEMON_API_LIST_URL)
	if err != nil {
		handleApiErrors(w, http.StatusInternalServerError, "")
		return
	}
	q := endpoint.Query()
	q.Set("limit", limit)
	q.Set("offset", offset)
	endpoint.RawQuery = q.Encode()

	resp, err := http.Get(endpoint.String())
	if err != nil {
		handleApiErrors(w, resp.StatusCode, "")
		return
	}

	var pokemonResponseResult PokemonApiListResponse
	jsonErr := json.NewDecoder(resp.Body).Decode(&pokemonResponseResult)
	if jsonErr != nil {
		handleApiErrors(w, http.StatusInternalServerError, "")
		return
	}

	var pokemonList []Pokemon

	wg := sync.WaitGroup{}

	for _, p := range pokemonResponseResult.Results {
		wg.Add(1)
		re := regexp.MustCompile(`/(\d+)`)
		id := strings.Replace(re.FindAllString(p.Url, 1)[0], "/", "", 1)

		go func(id string) {
			pokemon := getPokemonDetail(id)
			pokemonList = append(pokemonList, pokemon)
			wg.Done()
		}(id)
	}

	wg.Wait()

	if err != nil {
		handleApiErrors(w, http.StatusInternalServerError, "")
		return
	}
	jsonResponse, _ := json.Marshal(pokemonList)
	w.Write(jsonResponse)
}

func pokemonList(w http.ResponseWriter, req *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pokemonCollection := db.PokemonCollection()
	cursor, err := pokemonCollection.Find(
		ctx,
		bson.M{},
		// bson.M{"height": bson.D{{"$gt", 5}}}
	)
	if err != nil {
		panic(err)
	}

	var pokemons []Pokemon
	if err = cursor.All(ctx, &pokemons); err != nil {
		panic(err)
	}

	jsonResponse, _ := json.Marshal(pokemons)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonResponse)

}

func pokemonDetail(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, ok := vars["id"]
	if !ok {
		fmt.Println("id is missing in parameters")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	intId, _ := strconv.Atoi(id)

	var pokemon Pokemon
	pokemonCollection := db.PokemonCollection()
	pokemonCollection.FindOne(ctx, bson.M{"id": intId}).Decode(&pokemon)

	w.Header().Add("Content-Type", "application/json")
	if pokemon != (Pokemon{}) {
		jsonResponse, _ := json.Marshal(pokemon)
		w.Write(jsonResponse)
	} else {
		jsonResponse, _ := json.Marshal([]Pokemon{})
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonResponse)
	}

}

func main() {
	fmt.Println("hello from goapp")

	r := mux.NewRouter()
	r.HandleFunc("/goapp/pokemon", pokemonList)
	r.HandleFunc("/goapp/pokemon/{id}", pokemonDetail)
	r.HandleFunc("/goapp/api/pokemon", pokemonApiList)
	r.HandleFunc("/goapp/api/pokemon/{id}", pokemonApiDetail)

	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
