package server

import (
	"fmt"
	"log"
	"net/http"
)

// Init aaaaaaaaaaaaaaaaaaaaa
func Init() {
	fmt.Println("Starting a server ..")
	http.HandleFunc("/", index)
	//http.HandleFunc("/api/echo", api.EchoHandleFunc)
	//http.HandleFunc("/api/hello", api.HelloHandleFunc)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Fatal error, API server NOT started ! %v", err)
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to Rumia air monitoring system.")
}
