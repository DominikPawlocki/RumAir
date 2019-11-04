package server

import (	
	"github.com/robfig/cron/v3"
)

func ServerInit(){
	
	http.HandleFunc("/", index)
	//http.HandleFunc("/api/echo", api.EchoHandleFunc)
	//http.HandleFunc("/api/hello", api.HelloHandleFunc)

	http.ListenAndServe(port(), nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to Rumia air monitoring system.")
}
