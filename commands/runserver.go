package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to this life-changing API.\nIts the best API, its true, all other API's are fake.")
	})

	fmt.Println("Server listening!")
	http.ListenAndServe(":80", r)
}
