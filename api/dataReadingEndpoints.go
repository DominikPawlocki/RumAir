package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dompaw/RumAir_Pmpro_Sensors_API/airStations"
	"github.com/gorilla/mux"
)

// /stations/{stationId}/data?day=19&month=04year=2020&sensorCodes=blue|green|red dataGetting aaaa

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
//   description: The day of the reading(s). There will be 24 readings returned per every station's sensor, in 1h interval. If a sensorCodes are given in a query, only this one given will be returned.
//   required: true
//   type: string
//   format: should be like dd-MM-yyyy format
// - name: sensorCodes
//   in: query
//   description: The station sensors names. If a sensor doesnt belong to the station, will be ommited. There will be 24 readings returned per every sensor per given date( in 1 h intervals)
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

	vars := mux.Vars(r)
	stationID := vars["stationId"]
	sensorsQueryString := vars["sensorCodes"]
	year := vars["year"]
	month := vars["month"]
	day := vars["day"]

	var date time.Time
	var err error

	//------------check stationId - doesitExiist

	if date, err = dayMonthYearParser(day, month, year); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(fmt.Sprintf("stationId: %s, date: %v, sensors: %v", stationID, date, sensorsQueryString))

	begin, end := getBeginAndEndofTheDayInUnixEpoch(date)

	var sensors = getSensorsFromQueryStringBelongingToStationOnly(sensorsQueryString, stationID)
	if len(sensors) == 0 {
		http.Error(w, fmt.Sprintf("All the sensors from query : %s invalid", sensorsQueryString), http.StatusBadRequest)
		return
	}

	if _, err := airStations.GetSensorsDataBetweenTimePointsVIATIME(f, begin, end, airStations.A1H, sensors); err != nil {
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

func dayMonthYearParser(day string, month string, year string) (result time.Time, err error) {
	var sb strings.Builder
	var dayInt, monthInt, yearInt int

	if dayInt, err = strconv.Atoi(day); err != nil {
		sb.WriteString("Bad DAY argument. Seems not a number. /n")
	}
	if dayInt < 1 || dayInt > 31 {
		sb.WriteString("Day is to be between 1 or 31. ")
	}
	if monthInt, err = strconv.Atoi(month); err != nil {
		sb.WriteString("Bad MONTH argument. Seems not a number. /n")
	}
	if monthInt < 1 || monthInt > 12 {
		sb.WriteString("Month is to be between 1 or 12. ")
	}

	if yearInt, err = strconv.Atoi(year); err != nil {
		sb.WriteString("Bad YEAR argument. Seems not a number")
	}
	if yearInt < 2010 || yearInt > 2050 {
		sb.WriteString("YEAR is to be between 2010 or 2050. ")
	}
	if sb.Len() != 0 {
		err = errors.New(sb.String())
		return
	}
	result = time.Date(yearInt, time.Month(monthInt), dayInt, 0, 0, 0, 0, time.UTC)
	return
}

func getBeginAndEndofTheDayInUnixEpoch(date time.Time) (dayBegin int64, dayEnd int64) {
	year, month, day := date.Date()
	dayBegin = time.Date(year, month, day, 0, 0, 0, 0, time.UTC).Unix()
	dayEnd = time.Date(year, month, day, 23, 59, 59, 999, time.UTC).Unix()
	return
}

func getSensorsFromQueryStringBelongingToStationOnly(sensorsQueryString string, stationID string) (result []string) {
	querySplitted := strings.Split(sensorsQueryString, ",")
	result = []string{}

	for _, sensor := range querySplitted {
		//a trivial, initial check - let it be at least first two ASCII signs convertible to number, which is a stion naumber a sensor belongs to, in fact
		//stationNr, err := stationstrconv.ParseInt(sensor[1:2],10,8); err != nil{
		if strings.HasPrefix(sensor, stationID) {
			result = append(result, sensor)
		}
	}
	return
}
