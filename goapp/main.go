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

type OfficialArtwork struct {
	FrontDefault string `bson:"front_default" json:"front_default"`
}
type Other struct {
	OfficialArtwork OfficialArtwork `bson:"official-artwork" json:"official-artwork"`
}

type Sprites struct {
	Other Other `bson:"other" json:"other"`
}
type Pokemon struct {
	Id        int32     `bson:"id" json:"id"`
	Name      string    `bson:"name" json:"name"`
	Height    int32     `bson:"height,omitempty" json:"height"`
	Weight    int32     `bson:"weight,omitempty" json:"weight"`
	DateAdded time.Time `bson:"date_added" json:"date_added"`
	Sprites   Sprites   `bson:"sprites" json:"sprites"`
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

	jsonResponse, _ := json.Marshal(pokemon)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonResponse)

}

func main() {
	fmt.Println("hello from goapp")

	r := mux.NewRouter()
	r.HandleFunc("/goapp/pokemon", pokemonList)
	r.HandleFunc("/goapp/pokemon/{id}", pokemonDetail)

	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
