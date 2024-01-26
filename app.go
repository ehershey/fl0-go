package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/davecgh/go-spew/spew"
)

//go:embed static/*
var assets embed.FS

func main() {
	port := "3000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	fmt.Printf("port: %v\n", port)
	// A one-time verification occurs whenever creating or updating a subscription. Your server must respond to a request made to the GET method of your callback URL. The GET method must respond with a challenge code within 2 seconds. The request schema is as follows: https://{your-callback-url}?verification_token={request.verification_token}&challenge={random-string}. You should verify the verification_token is correct to ensure Oura is the one calling your API. Parse the challenge string from the query parameters and return the string in the body of your response. Example response body:

	// {
	// "challenge": "give-me-a-challenge"
	// }
	log.Println("setting things up")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		var querymap map[string][]string

		if r.Method == "GET" {
			querymap = r.URL.Query()
		} else {
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		}
		log.Println("got query")
		spew.Dump(querymap)
		// challenge = querymap["
		challenges, ok := r.URL.Query()["challenge"]
		if !ok || len(challenges) != 1 {
			http.Error(w, "challenge missing", http.StatusBadRequest)
			log.Println("challenge missing")
			return
		}
		challenge := challenges[0]
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "{\"challenge\":\"%s\"}\n", challenge)

		// err := t.Execute(w, "")
		// if err != nil {
		// fmt.Printf("got error: %v\n", err)
		// }
	})

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("got error: %v\n", err)
	}
}
