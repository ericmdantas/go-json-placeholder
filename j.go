package main

import (
	"net/http"
	"os"
)

func main() {
	var port string

	if port = os.Getenv("PORT"); port == "" {
		port = ":3333"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("yo!"))
	})

	http.ListenAndServe(port, nil)
}
