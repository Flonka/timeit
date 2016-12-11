package main

import (
	"encoding/json"
	"log"
	"net/http"
	"timeit"
)

type newEntryRequest struct {
	Name string `json:"name"`
}
type entryListResponse struct {
	Entries []timeit.Entry `json:"entries"`
}

func main() {

	http.HandleFunc("/api/entry/new", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var newReq newEntryRequest
		err := json.NewDecoder(r.Body).Decode(newReq)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Fatalln(err)
			return
		}

		log.Println("new entry: ", newReq)
		w.WriteHeader(http.StatusOK)

	})

	http.HandleFunc("/api/entry/list", getentries)

	address := ":8080"
	log.Println("Listening on " + address)
	log.Fatal(http.ListenAndServe(address, nil))
}

func getentries(w http.ResponseWriter, r *http.Request) {

	N := 10

	entries := make([]timeit.Entry, 0, N)

	for i := 0; i < N; i++ {
		e := timeit.NewEntry()
		e.Name = "Testing testing"
		entries = append(entries, *e)
	}

	var response entryListResponse
	response.Entries = entries
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
