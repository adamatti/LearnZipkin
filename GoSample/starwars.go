package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	mux "github.com/gorilla/mux"
)

func buildRequest(id string) *http.Request {
	url := "https://swapi.dev/api/people/" + id + "/"

	/*
		resp, err := resty.R().
			SetHeader("Content-Type", "application/json").
			Get(url)
	*/

	log.Println("Calling url:", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("unable to create http request: %+v\n", err)
	}

	return req
}

func callStarWars(zipkinObjects ZipkinObjects) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		starWarsRequest := buildRequest(vars["id"])

		resp, err := zipkinObjects.client.DoWithAppSpan(starWarsRequest, "some_function")
		defer resp.Body.Close()

		if err != nil {
			fmt.Fprintf(w, "There is an error: %+v", err)
		} else {
			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Fprint(w, string(body))
		}
	}
}
