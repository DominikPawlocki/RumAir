package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

// Hello response structure
type Hello struct {
	Message string
}

func HelloHandleFunc(w http.ResponseWriter, r *http.Request) {

	m := Hello{"Welcome to Rumia air monitoring system."}
	b, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to Rumia air monitoring system.")
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    key := vars["id"]

    fmt.Fprintf(w, "Key: " + key)
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // return the string response containing the request body    
    reqBody, _ := ioutil.ReadAll(r.Body)
    fmt.Fprintf(w, "%+v", string(reqBody))
}

/*resp := map[string]interface{}{
	"ok":      true,
	"balance": prevBalance + req.Amount,
}*/