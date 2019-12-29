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
				w.WriteHeader(500)
			} else {
				//w.WriteHeader(200)is called automatically
				w.Header().Add("Content-Type", "application/json; charset=utf-8")
				w.Write(resultBytes)
			}
		} else {
			resultBytes, _ = json.Marshal("Empty result Sensors per stations, when stations seems fetched ")
			w.WriteHeader(500)
		}
	} else {
		resultBytes, _ = json.Marshal("No stations can be fetched ")
	}

}
