package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dompaw/RumAir_Pmpro_Sensors_API/airStations"
)

// /stations/{stationId}/data?date=19042020&sensorCodes=blue|green|red dataGetting aaaa

// swagger:operation GET /stations/{stationId}/data dataGetting aaaa
// Gets a list of stations geolocalized (has CitiesNearby and Lat/Lon), with a nearby cities, using 3rd part service called LocationIQ.
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
// - name: date
//   in: query
//   description: The day of the reading(s). There will be 24 readings returned per every sensor, in 1h interval
//   required: true
//   type: string
//   format: should be like ddMMyyyy format
// - name: sensorCodes
//   in: query
//   description: The number of items to skip before starting to collect the result set
//   required: true
//   type: array
//   maxItems: 15
//   minItems: 1
//   unique: true
//   collectionFormat: csv
//   examples:
//     oneId:
//       summary: Example of a single ID
//       value: [5]   # ?ids=5
//     multipleIds:
//       summary: Example of multiple IDs
//       value: [1, 5, 7]   # ?ids=1,5,7
//   items:
//     type: string
//     minLength: 3
//     maxLength: 25
//     pattern: "\\w+"
// responses:
//   "200":
//     "$ref": "#/responses/geolocatingStationsResponse"
//   "404":
//     "$ref": "#/responses/notFound"
//   "500":
//     "$ref": "#/responses/internalServerError"
func AAAAAAAAAAAAA(w http.ResponseWriter, r *http.Request, f airStations.IHttpAbstracter) {
	var resultBytes []byte

	if _, err := airStations.GetSensorsDataBetweenTimePointsVIATIME(f, time.Now(), time.Now(), "someTimeOfAverage", []string{"04_blabla"}); err != nil {
		http.Error(w, fmt.Sprintf("%s %v", stationsCapabilitesFetchingError, err.Error()), http.StatusInternalServerError)
		return
	}
	// } else if localized, err := geolocalize.LocalizeStationsLocIQ(result); err != nil {
	// 	http.Error(w, fmt.Sprintf("%s %v", locationIQfetchingError, err.Error()), http.StatusInternalServerError)
	// 	return
	// } else if localized != nil {
	// 	result := geolocalize.LocalizedAirStationsResponse{
	// 		Localized:              localized,
	// 		WereLocalizedCount:     len(localized),
	// 		NotAbleToLocalizeCount: len(result) - len(localized)}

	// 	if resultBytes, err = json.Marshal(result); err != nil {
	// 		http.Error(w, fmt.Sprintf("%s %v", locationIQdeserializingError, err.Error()), http.StatusInternalServerError)
	// 		return
	// 	}
	// }
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(resultBytes)
}
