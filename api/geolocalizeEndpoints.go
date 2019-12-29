package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dompaw/RumAir/airStations"
	geolocalize "github.com/dompaw/RumAir/geoLocalize"
	"github.com/gorilla/mux"
)

/*myRouter.HandleFunc("/stations/locate/locationIQ", geolocalize.LocalizeStationsLocIQ)
myRouter.HandleFunc("/stations/{id}/locate/locationIQ", geolocalize.LocalizeStationsLocIQ)
myRouter.HandleFunc("/stations/locate/geobytes", geolocalize.LocalizeStationsGeoBytes(sts))*/

// Hello response structure
type Hello struct {
	Message string
}

func LocalizeAllStationsUsingLocationIQHandler(w http.ResponseWriter, r *http.Request) {
	var resultBytes []byte

	if result, err := airStations.GetAllStationsCapabilities(); err != nil {
		resultBytes, _ = json.Marshal(Hello{"Welcome to Rumia air monitoring system."})
	} else if localized, err := geolocalize.LocalizeStationsLocIQ(result); err != nil {
		resultBytes, _ = json.Marshal(Hello{err.Error()})
	} else if localized != nil {
		if resultBytes, err = json.Marshal(localized); err != nil {
			resultBytes, _ = json.Marshal(Hello{"Error during serializing response from geolocalizing stations via LocationIQ API."})
		}
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(resultBytes)
}

func LocalizeStationUsingLocationIQHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	stationID := vars["id"]

	fmt.Fprintf(w, "Key: "+stationID)

	var resultBytes []byte

	if result := airStations.GetStationCapabilities(stationID); len(result.Sensors) <= 0 {
		resultBytes, _ = json.Marshal(Hello{fmt.Sprintf("Didnt find any sensor for station id %v .", stationID)})
	} else if localized, err := geolocalize.LocalizeStationLocIQ(result); err != nil {
		resultBytes, _ = json.Marshal(Hello{err.Error()})
	} else if localized != nil {
		if resultBytes, err = json.Marshal(localized); err != nil {
			resultBytes, _ = json.Marshal(Hello{"Error during serializing response from geolocalizing stations via LocationIQ API."})
		}
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(resultBytes)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to Rumia air monitoring system.")
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
