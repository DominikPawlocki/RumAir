package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/dompaw/RumAir_Pmpro_Sensors_API/airStations"
)

// swagger:operation GET /stations/sensors/codes stationsAndSensors allSensorsCodesStringified
// Gets a list of sensors per station, stringified. Describes briefly which capability (sensors) has the particular station.
// ---
// produces:
// - application/json
// responses:
//   "200":
//     "$ref": "#/responses/allStationsSensorCodesStringifiedResponse"
//   "204":
//     "description": "It seems none of the stations has sensors. Probably some error occured ?"
//   "500":
//     "$ref": "#/responses/internalServerError"
func ShowAllStationsSensorCodesHandler(w http.ResponseWriter, r *http.Request, f airStations.IHttpAbstracter) {
	var resultBytes []byte

	if result, err := airStations.GetAllStationsCapabilities(f); err != nil {
		http.Error(w, fmt.Sprintf("%s %v", stationsCapabilitesFetchingError, err.Error()), http.StatusInternalServerError)
		return
	} else if len(result) > 0 {
		if sensorsPerStation := airStations.ShowStationsSensorsCodes(result); len(sensorsPerStation) > 0 {
			if resultBytes, err = json.Marshal(sensorsPerStation); err != nil {
				http.Error(w, fmt.Sprintf("%s %v", deserializingSensorsPerStationError, err.Error()), http.StatusInternalServerError)
				return
			}
			//w.WriteHeader(200)is called automatically
			w.Header().Add("Content-Type", "application/json; charset=utf-8")
			w.Write(resultBytes)

		} else {
			http.Error(w, fmt.Sprintf("%s %v", emptySensorsPerStationError, err.Error()), http.StatusNoContent)
			return
		}
	} else {
		http.Error(w, fmt.Sprintf("%s", emptySensorsPerStationError), http.StatusInternalServerError)
		return
	}
}

// swagger:operation GET /stations/sensors stationsAndSensors stationsFetching
// Gets a list of all system's (air)stations, with all its sensors (simplified model).
// ---
// produces:
// - application/json
// responses:
//   "200":
//     "$ref": "#/responses/stationsResponse"
//   "500":
//     "$ref": "#/responses/internalServerError"
func GetAllStationsCapabilitiesHandler(w http.ResponseWriter, r *http.Request, f airStations.IHttpAbstracter) {
	var resultBytes []byte
	if result, err := airStations.GetAllStationsCapabilities(f); err != nil {
		http.Error(w, fmt.Sprintf("%s %v", stationsCapabilitesFetchingError, err.Error()), http.StatusInternalServerError)
		return
	} else if resultBytes, err = json.Marshal(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(resultBytes)
	}
}

// swagger:operation GET /stations/{stationId}/sensors stationsAndSensors fullSensorsFetching
// Gets a list of sensors belonging to given station, with all the (sensors) properties (extended model).
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
//     "$ref": "#/responses/sensorsResponse"
//   "404":
//     "$ref": "#/responses/notFound"
//   "500":
//     "$ref": "#/responses/internalServerError"
func GetStationSensorsHandler(w http.ResponseWriter, r *http.Request, f airStations.IHttpAbstracter) {
	var resultBytes []byte

	//stationID := r.URL.Query()["message"][0]
	vars := mux.Vars(r)
	stationID := vars["id"]

	if result, err := airStations.GetStationSensors(f, stationID); err != nil {
		http.Error(w, fmt.Sprintf("%s %v", stationsCapabilitesFetchingError, err.Error()), http.StatusNotFound)
		return
	} else if resultBytes, err = json.Marshal(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(resultBytes)
	}
}

// swagger:operation GET /stations/{stationId}/sensors/codes stationsAndSensors sensorCodesOnlyFetching
// Returns only codes of all the sensors given station has.
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
//     "$ref": "#/responses/stationSensorCodesHandlerSuccessResp"
//   "404":
//     "$ref": "#/responses/notFound"
//   "500":
//     "$ref": "#/responses/internalServerError"
func GetStationSensorsOnlyCodesHandler(w http.ResponseWriter, r *http.Request, f airStations.IHttpAbstracter) {
	var resultBytes []byte

	vars := mux.Vars(r)
	stationID := vars["id"]

	if result, err := airStations.GetStationSensorCodesOnly(f, stationID); err != nil {
		http.Error(w, fmt.Sprintf("%s %v", stationsCapabilitesFetchingError, err.Error()), http.StatusNotFound)
		return
	} else if resultBytes, err = json.Marshal(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(resultBytes)
	}
}

// swagger:operation GET /stations/{stationId}/sensors/startTimes stationsAndSensors sensorsStartTimeFetching
// Returns the station sensors startTime. Shows which sensor began to work when.
// Unfortunately, the property 'StartDate' from sensor data is bullshit sometimes.. Eg station 23 - tells some sensor started to run in 2217, when in reality it started to collect data at June 2018...
// Have to find a better way to get that date of sensor start, then.
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
//     "$ref": "#/responses/stationSensorCodesHandlerSuccessResp"
//   "404":
//     "$ref": "#/responses/notFound"
//   "500":
//     "$ref": "#/responses/internalServerError"
func GetStationSensorsStartDatesHandler(w http.ResponseWriter, r *http.Request, f airStations.IHttpAbstracter) {
	var resultBytes []byte

	vars := mux.Vars(r)
	stationID := vars["id"]

	if result, err := airStations.GetSensorStartTimeAndCode(f, stationID); err != nil {
		http.Error(w, fmt.Sprintf("%s %v", stationsCapabilitesFetchingError, err.Error()), http.StatusNotFound)
		return
	} else if resultBytes, err = json.Marshal(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(resultBytes)
	}
}

/*
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
