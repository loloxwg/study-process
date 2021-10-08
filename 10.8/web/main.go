package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("hello from snippets"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d\n", id)
	//w.Write([]byte("Display a special snippet..."))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {

		w.Header().Set("Allow", "POST")
		w.Header()["Date"] = nil
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"name":"Alex"}`))
		http.Error(w, "Method not found", 405)
		//w.WriteHeader(405)
		//w.Write([]byte("Method not allowed"))
		return
	}
	w.Write([]byte("Create a new snippet"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
	log.Println("Starting http server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
