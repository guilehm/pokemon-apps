package main

import (
	"fmt"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")

}

func main() {
	fmt.Println("Hello World")
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
