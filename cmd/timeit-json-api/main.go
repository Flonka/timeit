package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type newEntryRequest struct {
	Name string `json:"name"`
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

	address := ":8080"
	log.Println("Listening on " + address)
	log.Fatal(http.ListenAndServe(address, nil))
}
