package main

import (
	"context"
	"encoding/json"
	"fmt"
	"goapp/db"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
)

const POKEMON_API_BASE_URL = "https://pokeapi.co/api/v2"
const POKEMON_API_DETAIL_URL = "https://pokeapi.co/api/v2/pokemon/"

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

func pokemonApiDetail(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, ok := vars["id"]
	if !ok {
		fmt.Println("id is missing in parameters")
	}

	endpoint := POKEMON_API_DETAIL_URL + id
	resp, err := http.Get(endpoint)

	if err != nil {
		panic(err)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		panic(readErr)
	}

	pokemon := Pokemon{}

	jsonErr := json.Unmarshal(body, &pokemon)
	if jsonErr != nil {
		panic(jsonErr)
	}

	jsonResponse, err := json.Marshal(pokemon)

	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
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
	r.HandleFunc("/goapp/api/pokemon/{id}", pokemonApiDetail)

	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
