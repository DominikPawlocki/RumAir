package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dompaw/RumAir/airStations"
	geolocalize "github.com/dompaw/RumAir/geoLocalize"
	"github.com/gorilla/mux"
)

// .../stations/locate/geobytes
func LocalizeAllStationsUsingGeoBytesHandler(w http.ResponseWriter, r *http.Request) {
	var resultBytes []byte

	if result, err := airStations.GetAllStationsCapabilities(airStations.StationsCapabiltiesFetcher{}); err != nil {
		http.Error(w, fmt.Sprintf("%s %v", stationsCapabilitesFetchingError, err.Error()), http.StatusInternalServerError)
		return
	} else if localized, err := geolocalize.LocalizeStationsGeoBytes(result); err != nil {
		http.Error(w, fmt.Sprintf("%s %v", geoBytesfetchingError, err.Error()), http.StatusInternalServerError)
		return
	} else if localized != nil {
		if resultBytes, err = json.Marshal(localized); err != nil {
			http.Error(w, fmt.Sprintf("%s %v", geoBytesdeserializingError, err.Error()), http.StatusInternalServerError)
			return
		}
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(resultBytes)
}

// .../stations/locate/locationIQ)
func LocalizeAllStationsUsingLocationIQHandler(w http.ResponseWriter, r *http.Request) {
	var resultBytes []byte

	if result, err := airStations.GetAllStationsCapabilities(airStations.StationsCapabiltiesFetcher{}); err != nil {
		http.Error(w, fmt.Sprintf("%s %v", stationsCapabilitesFetchingError, err.Error()), http.StatusInternalServerError)
		return
	} else if localized, err := geolocalize.LocalizeStationsLocIQ(result); err != nil {
		http.Error(w, fmt.Sprintf("%s %v", locationIQfetchingError, err.Error()), http.StatusInternalServerError)
		return
	} else if localized != nil {
		if resultBytes, err = json.Marshal(localized); err != nil {
			http.Error(w, fmt.Sprintf("%s %v", locationIQdeserializingError, err.Error()), http.StatusInternalServerError)
			return
		}
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(resultBytes)
}

// .../stations/{id}/locate/locationIQ"
func LocalizeStationUsingLocationIQHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	stationID := vars["id"]

	var resultBytes []byte

	if result := airStations.GetStationCapabilities(airStations.StationsCapabiltiesFetcher{}, stationID); len(result.Sensors) <= 0 {
		http.Error(w, fmt.Sprintf("Cannot fectch station with ID %s", stationID), http.StatusInternalServerError)
		return
	} else if localized, err := geolocalize.LocalizeStationLocIQ(result); err != nil {
		http.Error(w, fmt.Sprintf("%s %v", locationIQfetchingError, err.Error()), http.StatusInternalServerError)
		return
	} else if localized != nil {
		if resultBytes, err = json.Marshal(localized); err != nil {
			http.Error(w, fmt.Sprintf("%s %v", locationIQdeserializingError, err.Error()), http.StatusInternalServerError)
		}
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(resultBytes)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to Rumia air monitoring system.")
}
