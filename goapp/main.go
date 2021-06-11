package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var CNX = Connection()

func Connection() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		panic(err)
	}
	return client
}

type Pokemon struct {
	Id        int32     `bson:"id"`
	Name      string    `bson:"name"`
	Height    int32     `bson:"height,omitempty"`
	Weight    int32     `bson:"weight,omitempty"`
	DateAdded time.Time `bson:"dateAdded"`
}

func pokemonList(w http.ResponseWriter, req *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client := Connection()

	database := client.Database("spoon")
	pokemonCollection := database.Collection("pokemons")

	var pokemons []Pokemon
	cursor, err := pokemonCollection.Find(ctx, bson.M{"height": bson.D{{"$gt", 5}}})
	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &pokemons); err != nil {
		panic(err)
	}

	jsonResponse, _ := json.Marshal(pokemons)
	w.Write(jsonResponse)

}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")

}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func pokemonDetail(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, ok := vars["id"]
	if !ok {
		fmt.Println("id is missing in parameters")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := Connection()
	database := client.Database("spoon")
	pokemonCollection := database.Collection("pokemons")

	var pokemon Pokemon

	intId, _ := strconv.Atoi(id)

	pokemonCollection.FindOne(ctx, bson.M{"id": intId}).Decode(&pokemon)

	jsonResponse, _ := json.Marshal(pokemon)
	w.Write(jsonResponse)

}

func main() {
	fmt.Println("Hello World")

	r := mux.NewRouter()
	r.HandleFunc("/goapp/hello", hello)
	r.HandleFunc("/goapp/headers", headers)
	r.HandleFunc("/goapp/pokemon", pokemonList)
	r.HandleFunc("/goapp/pokemon/{id}", pokemonDetail)

	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
