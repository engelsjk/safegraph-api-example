package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type graphQLRequest struct {
	Query     string `json:"query"`
	Variables string `json:"variables,omitempty"`
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	query := `query($placekeys: [Placekey!]) {
		places(placekeys: $placekeys) {
			placekey
			safegraph_core {
				location_name
				top_category
				street_address
				city
				region
				latitude
				longitude
			}
		}
	}`

	variables := `{
		"placekeys": [
			"222-222@5qw-shj-7qz",
			"222-222@5s6-pyc-7qz",
			"zzw-222@5vg-3tv-7qz",
			"22s-223@63j-4f8-7qz",
			"228-222@5sb-cyn-7qz"
		]
	}`

	gql, err := json.Marshal(graphQLRequest{Query: query, Variables: variables})
	if err != nil {
		panic(err)
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://api.safegraph.com/v1/graphql", bytes.NewBuffer(gql))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("apikey", os.Getenv("SAFEGRAPH_API_KEY"))
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	rb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(os.Getenv("SAFEGRAPH_API_KEY"))
	log.Println(string(gql))
	log.Println(resp.Status)
	log.Println(resp.Header)
	log.Println(string(rb))
}
