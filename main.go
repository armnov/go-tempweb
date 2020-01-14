package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func homeController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homeController)
	fmt.Println("Running on port 8080")
	http.ListenAndServe(":8080", router)
}
