package main

import (
	"fmt"
	"log"
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string("gospock lives"))
}

func main() {
	fmt.Println("started gospock")
	http.HandleFunc("/health", health)
	err := http.ListenAndServe(":4500", nil)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("stopped gospock")
}
