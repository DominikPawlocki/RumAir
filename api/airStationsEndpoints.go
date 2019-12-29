package api

import (
	"encoding/json"
	"net/http"

	"github.com/dompaw/RumAir/airStations"
)

func ShowStationsSensorsCodesHandler(w http.ResponseWriter, r *http.Request) {
	var resultBytes []byte

	if result := airStations.GetAllStationsCapabilities(); len(result) > 0 {
		if sensorsPerStation := airStations.ShowStationsSensorsCodes(result); len(sensorsPerStation) > 0 {
			resultBytes, _ = json.Marshal(sensorsPerStation)
		} else {
			resultBytes, _ = json.Marshal("Empty result Sensors per stations, when stations seems fetched ")
		}
	} else {
		resultBytes, _ = json.Marshal("No stations can be fetched ")
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(resultBytes)
}
