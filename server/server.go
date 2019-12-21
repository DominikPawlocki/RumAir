package server

// this is just a dummy API strarting point. Will be coded properly later/moved to API.

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Init() {
	fmt.Println("Starting a server ..")
	http.HandleFunc("/", index)
	//http.HandleFunc("/api/echo", api.EchoHandleFunc)
	http.HandleFunc("/api/hello", HelloHandleFunc)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Fatal error, API server NOT started ! %v", err)
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to Rumia air monitoring system.")
}

// Hello response structure
type Hello struct {
	Message string
}

// Should be moved to API
func HelloHandleFunc(w http.ResponseWriter, r *http.Request) {

	m := Hello{"Welcome to Rumia air monitoring system."}
	b, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}

/*resp := map[string]interface{}{
	"ok":      true,
	"balance": prevBalance + req.Amount,
}*/
