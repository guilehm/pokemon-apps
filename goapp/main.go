package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Pokemon struct {
	Id        int32     `bson:"id"`
	Name      string    `bson:"name"`
	Height    int32     `bson:"height,omitempty"`
	Weight    int32     `bson:"weight,omitempty"`
	DateAdded time.Time `bson:"dateAdded"`
}

func pokemon(w http.ResponseWriter, req *http.Request) {

	uri := os.Getenv("MONGODB_URI")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to Mongodb and pinged.")

	http.HandleFunc("/goapp/hello", hello)
	http.HandleFunc("/goapp/headers", headers)

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
