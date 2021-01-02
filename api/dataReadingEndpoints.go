package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dompaw/RumAir_Pmpro_Sensors_API/airStations"
	"github.com/gorilla/mux"
)

func GetSingleDayOfStationSensorsReadings(w http.ResponseWriter, r *http.Request, f airStations.IHttpAbstracter) {
	// swagger:operation GET /stations/{stationId}/data dataGetting dailyDataFetching
	// Gets a sensor's data readings, starting at a 00:00:00 of given day.
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: stationId
	//   in: path
	//   description: weather station id
	//   required: true
	//   type: string
	//   format: like 02, 04 etc
	// - name: day
	//   in: query
	//   description: The day part of the date
	//   required: true
	//   minimum: 1
	//   maximum: 31
	//   type: integer
	//   format: dd format
	// - name: month
	//   in: query
	//   description: The month part of the date (of reading(s)).
	//   required: true
	//   minimum: 1
	//   maximum: 12
	//   type: integer
	//   format: MM format
	// - name: year
	//   in: query
	//   description: The year part of the date (of reading(s)).
	//   required: true
	//   minimum: 2014
	//   maximum: 2050
	//   type: integer
	//   format: yyyy format
	// - name: interval
	//   in: query
	//   description: Means the average the (pmPro) system does for readings. A1h - hour averages, so by given day there will be 24 results returned per every sensor. A24h - means dayily averages, so there will be 1 result per day returned.
	//   required: true
	//   type: string
	//   enum: [A1h, A24h]
	// - name: sensorCodes
	//   in: query
	//   description: The station sensors names. If a sensor doesnt belong to the station, will be ommited. If this is not provided, the endpoint will return data from ALL the sensors the particular station has. It needs to fetch those (sensor) first.. Performance costly and much slower then !
	//   required: false
	//   type: array
	//   maxItems: 15
	//   minItems: 1
	//   unique: true
	//   collectionFormat: csv
	//   items:
	//     type: string
	//     minLength: 3
	//     maxLength: 25
	//     pattern: "\\w+"
	// responses:
	//   "200":
	//     "$ref": "#/responses/sensorDataKeyedViaCodeHandlerResponse"
	//   "400":
	//     "$ref": "#/responses/badRequest"
	//   "404":
	//     "$ref": "#/responses/notFound"
	//   "500":
	//     "$ref": "#/responses/internalServerError"
	vars := mux.Vars(r)
	stationID := vars["stationId"]
	interval := vars["interval"]
	sensorsQueryString := vars["sensorCodes"]
	year := vars["year"]
	month := vars["month"]
	day := vars["day"]

	var date time.Time
	var err error

	//--------- todo : check stationId - if it doesnt exist

	if date, err = dayMonthYearParser(day, month, year); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(fmt.Sprintf("stationId: %s, date: %v, sensors: %v", stationID, date, sensorsQueryString))

	switch interval {
	case "A1h":
		interval = airStations.A1H
	case "A24h":
		interval = airStations.A24H
	default:
		http.Error(w, fmt.Sprintf("%s not supported for now. 'A1h' and 'A24h' supported.", interval), http.StatusBadRequest)
		return
	}

	begin, end := getBeginAndEndofTheDayInUnixEpoch(date)

	resultBytes, err := runTheFlow(sensorsQueryString, stationID, begin, end, interval, f)

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(resultBytes)
}

func runTheFlow(sensorsQueryString string, stationID string, begin int64, end int64, timeofAverage string, f airStations.IHttpAbstracter) (resultBytes []byte, err error) {
	var sensors []string

	if sensors, err = fetchOrProcessSensorCodes(sensorsQueryString, stationID, f); err != nil {
		return
	}

	var sensorCodeKeyedResp = airStations.SensorDataKeyedViaCode{}
	if sensorCodeKeyedResp, _, err = airStations.GetSensorsDataBetweenTimePoints(f, begin, end, timeofAverage, sensors); err != nil {
		return nil, fmt.Errorf("%s %v", stationsCapabilitesFetchingError, err.Error())
	}
	if resultBytes, err = json.Marshal(sensorCodeKeyedResp); err != nil {
		return nil, fmt.Errorf("%s %v", locationIQdeserializingError, err.Error())
	}
	return
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

func filterQueryStringToGetSensorsCodesBelongingToStationOnly(sensorsQueryString string, stationID string) (result []string) {
	querySplitted := strings.Split(sensorsQueryString, ",")

	for _, sensor := range querySplitted {
		//a trivial, initial check - let it be at least first two ASCII signs convertible to number, which is a stion number a sensor belongs to, in fact
		//stationNr, err := stationstrconv.ParseInt(sensor[1:2],10,8); err != nil{
		if strings.HasPrefix(sensor, stationID) {
			result = append(result, sensor)
		}
	}
	return
}

func fetchOrProcessSensorCodes(sensorsQueryString string, stationID string, f airStations.IHttpAbstracter) (sensors []string, err error) {
	if sensorsQueryString != "" {
		sensors = filterQueryStringToGetSensorsCodesBelongingToStationOnly(sensorsQueryString, stationID)
		if len(sensors) == 0 {
			err = fmt.Errorf("All the sensors from query : %s invalid, or doesnt belong to given station.", sensorsQueryString)
			return
		}
	} else {
		//means taking ALL the station sensors. Performance expensive !
		if sensors, err = airStations.GetStationSensorCodesOnly(f, stationID); err != nil {
			return
		}
	}
	return
}
