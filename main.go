package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	fmt.Fprintf(w, "hai %s", name)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{name}", greetHandler).Methods("GET")

	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", r)
}
