package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dompaw/RumAir_Pmpro_Sensors_API/airStations"
	geolocalize "github.com/dompaw/RumAir_Pmpro_Sensors_API/geoLocalize"
	"github.com/gorilla/mux"
)

func LocalizeAllStationsUsingGeoBytesHandler(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /stations/locate/geobytes geolocating geolocatingByGeobytesResponse
	// Gets a list of stations geolocalized (has CitiesNearby and Lat/Lon), with a nearby cities, using 3rd part service called GeoBytes.
	// ---
	// produces:
	// - application/json
	// responses:
	//   "200":
	//     "$ref": "#/responses/geolocatingStationsResponse"
	//   "404":
	//     "$ref": "#/responses/notFound"
	//   "500":
	//     "$ref": "#/responses/internalServerError"
	var resultBytes []byte

	if result, err := airStations.GetAllStationsCapabilities(airStations.HttpAbstracter{}); err != nil {
		http.Error(w, fmt.Sprintf("%s %v", stationsCapabilitesFetchingError, err.Error()), http.StatusInternalServerError)
		return
	} else if localized, err := geolocalize.LocalizeStationsGeoBytes(result); err != nil {
		http.Error(w, fmt.Sprintf("%s %v", geoBytesfetchingError, err.Error()), http.StatusInternalServerError)
		return
	} else if localized != nil {
		result := geolocalize.LocalizedAirStationsResponse{
			Localized:              localized,
			WereLocalizedCount:     len(localized),
			NotAbleToLocalizeCount: len(result) - len(localized)}

		if resultBytes, err = json.Marshal(result); err != nil {
			http.Error(w, fmt.Sprintf("%s %v", geoBytesdeserializingError, err.Error()), http.StatusInternalServerError)
			return
		}
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(resultBytes)
}

func LocalizeAllStationsUsingLocationIQHandler(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /stations/locate/locationIQ geolocating geolocatingByLocIQResponse
	// Gets a list of stations geolocalized (has CitiesNearby and Lat/Lon), with a nearby cities, using 3rd part service called LocationIQ.
	// ---
	// produces:
	// - application/json
	// responses:
	//   "200":
	//     "$ref": "#/responses/geolocatingStationsResponse"
	//   "404":
	//     "$ref": "#/responses/notFound"
	//   "500":
	//     "$ref": "#/responses/internalServerError"
	var resultBytes []byte

	if result, err := airStations.GetAllStationsCapabilities(airStations.HttpAbstracter{}); err != nil {
		http.Error(w, fmt.Sprintf("%s %v", stationsCapabilitesFetchingError, err.Error()), http.StatusInternalServerError)
		return
	} else if localized, err := geolocalize.LocalizeStationsLocIQ(result); err != nil {
		http.Error(w, fmt.Sprintf("%s %v", locationIQfetchingError, err.Error()), http.StatusInternalServerError)
		return
	} else if localized != nil {
		result := geolocalize.LocalizedAirStationsResponse{
			Localized:              localized,
			WereLocalizedCount:     len(localized),
			NotAbleToLocalizeCount: len(result) - len(localized)}

		if resultBytes, err = json.Marshal(result); err != nil {
			http.Error(w, fmt.Sprintf("%s %v", locationIQdeserializingError, err.Error()), http.StatusInternalServerError)
			return
		}
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(resultBytes)
}

func LocalizeStationUsingLocationIQHandler(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /stations/{stationId}/locate/locationIQ geolocating geolocateParticularStation
	// Geolocalize (finds CitiesNearby and Lat/Lon) one particular station by 3rd party service 'LocationIQ', and returns it with its sensors simplified model also.
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: stationId
	//   in: path
	//   description: weather station id
	//   required: true
	//   type: string
	//   format: should be like 02, 04 etc
	// responses:
	//   "200":
	//     "$ref": "#/responses/geolocatingParticularStationResponse"
	//   "404":
	//     "$ref": "#/responses/notFound"
	//   "500":
	//     "$ref": "#/responses/internalServerError"
	vars := mux.Vars(r)
	stationID := vars["id"]

	var resultBytes []byte

	if result := airStations.GetStationCapabilities(airStations.HttpAbstracter{}, stationID); len(result.Sensors) <= 0 {
		http.Error(w, fmt.Sprintf("Cannot fetch station with ID %s", stationID), http.StatusNotFound)
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

func GetStationNumbersPerCityHandler(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /stations/locate/locationIQ/numbersPerCity geolocating geolocatingCitiesWithStationsResponse
	// Gets a list of geolocalized station IDs, groupped by nearby city(ies), using 3rd part service called LocationIQ.
	// ---
	// produces:
	// - application/json
	// responses:
	//   "200":
	//     "$ref": "#/responses/geolocatingCitiesWithStationsResponse"
	//   "404":
	//     "$ref": "#/responses/notFound"
	//   "500":
	//     "$ref": "#/responses/internalServerError"
	var resultBytes []byte

	if result, err := airStations.GetAllStationsCapabilities(airStations.HttpAbstracter{}); err != nil {
		http.Error(w, fmt.Sprintf("%s %v", stationsCapabilitesFetchingError, err.Error()), http.StatusInternalServerError)
		return
	} else if localized, err := geolocalize.LocalizeStationsLocIQ(result); err != nil {
		http.Error(w, fmt.Sprintf("%s %v", locationIQfetchingError, err.Error()), http.StatusInternalServerError)
		return
	} else if localized != nil {
		citiesWithStations := geolocalize.GetStationNrPerCity(localized)
		result := geolocalize.CitiesWithStations{
			Localized:              citiesWithStations,
			WereLocalizedCount:     len(localized),
			NotAbleToLocalizeCount: len(result) - len(localized)}

		if resultBytes, err = json.Marshal(result); err != nil {
			http.Error(w, fmt.Sprintf("%s %v", locationIQdeserializingError, err.Error()), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(resultBytes)
}
