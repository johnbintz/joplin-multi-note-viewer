package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")

	response, err := performSearch(query)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Issue querying Joplin! Is it running?"))
		return
	}

	body, err := response.GetFrontendBody()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error serializing JSON!"))
		return
	}

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Write(body)
}

func noteHandler(w http.ResponseWriter, r *http.Request) {
	noteID := mux.Vars(r)["id"]

	response, err := retrieveNote(noteID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Issue querying Joplin! Is it running?"))
		return
	}

	body, err := response.GetFrontendBody()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error serializing JSON!"))
		return
	}

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Write(body)
}
