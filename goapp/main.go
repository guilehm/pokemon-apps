package main

import (
	"context"
	"encoding/json"
	"fmt"
	"goapp/db"
	"goapp/rabbit"
	"goapp/utils"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/streadway/amqp"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const POKEMON_API_BASE_URL = "https://pokeapi.co/api/v2"
const POKEMON_API_LIST_URL = "https://pokeapi.co/api/v2/pokemon/"

type Pokemon struct {
	Id      int32  `bson:"id" json:"id"`
	Name    string `bson:"name" json:"name"`
	Height  int32  `bson:"height,omitempty" json:"height,omitempty"`
	Weight  int32  `bson:"weight,omitempty" json:"weight,omitempty"`
	Sprites struct {
		Other struct {
			OfficialArtwork struct {
				FrontDefault string `bson:"front_default" json:"front_default"`
			} `bson:"official-artwork" json:"official-artwork"`
		} `bson:"other" json:"other"`
	} `bson:"sprites" json:"sprites"`
	Types []struct {
		Slot int32 `bson:"slot" json:"slot"`
		Type struct {
			Name string `bson:"name" json:"name"`
			Url  string `bson:"url" json:"url"`
		} `bson:"type" json:"type"`
	} `bson:"types" json:"types"`
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

func getPokemonDetail(id string, ch *amqp.Channel, q amqp.Queue) Pokemon {
	endpoint := POKEMON_API_LIST_URL + id
	resp, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}
	var pokemon Pokemon
	body, _ := ioutil.ReadAll(resp.Body)
	if err = json.Unmarshal(body, &pokemon); err != nil {
		panic(err)
	}

	now := time.Now()
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
			MessageId:   fmt.Sprint(now.UnixNano()),
		},
	)

	rabbit.FailOnError(err, "Failed to publish a message")
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
	if err = json.NewDecoder(resp.Body).Decode(&pokemon); err != nil {
		handleApiErrors(w, http.StatusInternalServerError, "")
		return
	}

	jsonResponse, _ := json.Marshal(pokemon)
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

	conn, connError := rabbit.RabbitConnection()
	rabbit.FailOnError(connError, "Failed to connect to RabbitMQ")
	ch, err := conn.Channel()
	defer conn.Close()
	rabbit.FailOnError(err, "Failed to open a channel")
	defer ch.Close()
	queue, qError := ch.QueueDeclare(
		os.Getenv("POKEMON_ROUTING_KEY"),
		false,
		false,
		false,
		false,
		nil,
	)

	rabbit.FailOnError(qError, "Failed to declare a queue")

	pokemonList := make([]Pokemon, 0)
	wg := sync.WaitGroup{}

	for _, p := range pokemonResponseResult.Results {
		wg.Add(1)
		re := regexp.MustCompile(`/(\d+)`)
		id := strings.Replace(re.FindAllString(p.Url, 1)[0], "/", "", 1)

		go func(id string) {
			pokemon := getPokemonDetail(id, ch, queue)
			pokemonList = append(pokemonList, pokemon)
			wg.Done()
		}(id)
	}

	wg.Wait()

	sort.Slice(pokemonList, func(i, j int) bool {
		return pokemonList[i].Id < pokemonList[j].Id
	})

	jsonResponse, _ := json.Marshal(pokemonList)
	w.Write(jsonResponse)
}

func pokemonList(w http.ResponseWriter, req *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pokemonCollection := db.PokemonCollection()
	opts := options.Find().SetProjection(
		bson.M{"_id": 0, "id": 1, "name": 1, "sprites.other.official-artwork.front_default": 1},
	)
	cursor, err := pokemonCollection.Find(
		ctx,
		bson.M{},
		opts,
	)
	if err != nil {
		handleApiErrors(w, http.StatusInternalServerError, "")
		return
	}

	var pokemons []Pokemon
	if err = cursor.All(ctx, &pokemons); err != nil {
		handleApiErrors(w, http.StatusInternalServerError, "")
		return
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
	if pokemon.Name != "" {
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
	r.StrictSlash(true).HandleFunc("/goapp/pokemon/", pokemonList)
	r.StrictSlash(true).HandleFunc("/goapp/pokemon/{id}/", pokemonDetail)
	r.StrictSlash(true).HandleFunc("/goapp/api/pokemon/", pokemonApiList)
	r.StrictSlash(true).HandleFunc("/goapp/api/pokemon/{id}/", pokemonApiDetail)

	http.ListenAndServe(":"+os.Getenv("PORT"), utils.LogRequest(r))

}
