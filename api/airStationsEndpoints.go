package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dompaw/RumAir/airStations"
)

// ...stations/sensors/codes
func ShowAllStationsSensorsCodesHandler(w http.ResponseWriter, r *http.Request, f airStations.IStationsCapabiltiesFetcher) {
	var resultBytes []byte

	if result, err := airStations.GetAllStationsCapabilities(f); err != nil {
		http.Error(w, fmt.Sprintf("%s %v", stationsCapabilitesFetchingError, err.Error()), http.StatusInternalServerError)
		return
	} else if len(result) > 0 {
		if sensorsPerStation := airStations.ShowStationsSensorsCodes(result); len(sensorsPerStation) > 0 {
			if resultBytes, err = json.Marshal(sensorsPerStation); err != nil {
				http.Error(w, fmt.Sprintf("%s %v", deserializingSensorsPerStationError, err.Error()), http.StatusInternalServerError)
				return
			} else {
				//w.WriteHeader(200)is called automatically
				w.Header().Add("Content-Type", "application/json; charset=utf-8")
				w.Write(resultBytes)
			}
		} else {
			http.Error(w, fmt.Sprintf("%s %v", emptySensorsPerStationError, err.Error()), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, fmt.Sprintf("%s", emptySensorsPerStationError), http.StatusInternalServerError)
		return
	}

}

// .../stations/sensors
func GetAllStationsCapabilitiesHandler(w http.ResponseWriter, r *http.Request, f airStations.IStationsCapabiltiesFetcher) {
	var resultBytes []byte
	if result, err := airStations.GetAllStationsCapabilities(f); err != nil {
		http.Error(w, fmt.Sprintf("%s %v", stationsCapabilitesFetchingError, err.Error()), http.StatusInternalServerError)
		return
	} else if resultBytes, err = json.Marshal(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		//w.Write([]byte("Unsupported request method."))
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(resultBytes)
	}
}

/*
func EchoHandleFunc(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query()["message"][0]

	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, message)
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))
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
