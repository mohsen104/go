package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/google", http.RedirectHandler("http://goolge.com", 307))
	mux.HandleFunc("/users", GetUser)

	server := http.Server{
		Addr:         ":8080",
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		Handler:      mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(205)
	fmt.Fprintln(w, "Helloooo")
}
