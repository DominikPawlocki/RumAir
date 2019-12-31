package api

import (
	"encoding/json"
	"net/http"

	"github.com/dompaw/RumAir/airStations"
)

func ShowStationsSensorsCodesHandler(w http.ResponseWriter, r *http.Request) {
	var resultBytes []byte

	if result, err := airStations.GetAllStationsCapabilities(); err != nil {
		resultBytes, _ = json.Marshal(err.Error())
		w.WriteHeader(500)
	} else if len(result) > 0 {
		if sensorsPerStation := airStations.ShowStationsSensorsCodes(result); len(sensorsPerStation) > 0 {
			resultBytes, err = json.Marshal(sensorsPerStation)
			if err != nil {
				resultBytes, _ = json.Marshal("Error on deserializing sensorsPerStation, when stations seems fetched ")
				w.WriteHeader(500)
				w.Write(resultBytes)
			} else {
				//w.WriteHeader(200)is called automatically
				w.Header().Add("Content-Type", "application/json; charset=utf-8")
				w.Write(resultBytes)
			}
		} else {
			resultBytes, _ = json.Marshal("Empty result Sensors per stations, when stations seems fetched ")
			w.WriteHeader(500)
			w.Write(resultBytes)
		}
	} else {
		resultBytes, _ = json.Marshal("No stations can be fetched ")
		w.WriteHeader(500)
		w.Write(resultBytes)
	}

}

/*
func EchoHandleFunc(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query()["message"][0]

	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, message)
}


// Decode JSON
	user := &User{}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(user); err != nil {
		return nil, err
	}

		if created {
			w.Header().Add("Location", "/api/books/"+isbn)
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusConflict)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method."))
	}



	func sendResponse(entry *Entry, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	if err := enc.Encode(entry); err != nil {
		log.Printf("error encoding %+v - %s", entry, err)
	}



	func kvPostHandler(w http.ResponseWriter, r *http.Request) {
	// Decode request
		...
	dec := json.NewDecoder(r.Body)
	entry := &Entry{}
	if err := dec.Decode(entry); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !ok {
		http.Error(w, fmt.Sprintf("Key %q not found", key), http.StatusNotFound)
		return
	}


	/*resp := map[string]interface{}{
	"ok":      true,
	"balance": prevBalance + req.Amount,
*/
