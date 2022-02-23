package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"mime"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	//"github.com/pkg/browser"
)

//go:embed dist/*
var res embed.FS

func main() {
	mime.AddExtensionType("js", "application/javascript")
	mime.AddExtensionType("css", "text/css")

	err := readConfig()
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/v1/search", searchHandler)
	router.HandleFunc("/api/v1/notes", notesHandler)
	router.HandleFunc("/api/v1/note/{id}", noteHandler)

	fmt.Println("starting public file server")

	strippedFs, err := fs.Sub(res, "dist")
	if err != nil {
		log.Fatal("can't load embedded js")
	}

	fs := http.FileServer(http.FS(strippedFs))

	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, ".css") {
			w.Header().Set("Content-Type", "text/css")
		} else if strings.HasSuffix(r.URL.Path, ".js") {
			w.Header().Set("Content-Type", "application/javascript")
		}

		fs.ServeHTTP(w, r)
	})

	fmt.Println("listening on localhost:8181")
	http.ListenAndServe("localhost:8181", router)

	/*
		err = browser.OpenURL("http://localhost:8181")
		if err != nil {
			fmt.Println("Unable to open browser, visit http://localhost:8181")
		}
	*/
}
