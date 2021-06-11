package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Pokemon struct {
	Id        int32     `bson:"id"`
	Name      string    `bson:"name"`
	Height    int32     `bson:"height,omitempty"`
	Weight    int32     `bson:"weight,omitempty"`
	DateAdded time.Time `bson:"dateAdded"`
}

func pokemon(w http.ResponseWriter, req *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

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

	json, _ := json.Marshal(pokemons)
	w.Write(json)

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

func main() {
	fmt.Println("Hello World")

	http.HandleFunc("/goapp/hello", hello)
	http.HandleFunc("/goapp/pokemon", pokemon)
	http.HandleFunc("/goapp/headers", headers)

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
