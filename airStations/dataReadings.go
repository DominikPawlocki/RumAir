package airStations

import (
	"strconv"
	"time"
)

type SensorDataKeyedViaTime struct {
	Data []sensorDataKeyedViaTime `json:"data"`
}

type sensorDataKeyedViaTime struct {
	Time int64 `json:"time"`
	Data []struct {
		SensorCode string  `json:"sensorCode"`
		Value      float32 `json:"value"`
	} `json:"data"`
}

type SensorDataKeyedViaCode struct {
	Data []sensorDataKeyedViaCode `json:"data"`
}

type sensorDataKeyedViaCode struct {
	SensorCode string `json:"sensorCode"`
	Data       []struct {
		Time  int64   `json:"time"`
		Value float32 `json:"value"`
	} `json:"data"`
}

// used also in geoLocalize package. this is a responce from PmPro system, not response from this API. Probably should be organized better ..
type PmProSensorsDataInTimePeriodResponse struct {
	EndTime   int64 `json:"end"`
	StartTime int64 `json:"start"`
	Data      [][]struct {
		Time  int64   `json:"t"` // "t": 1597611600, "v": 62.780000000000000,
		Value float32 `json:"v"`
	} `json:"values"`
	Sensors []string `json:"vars"`
}

func GetSensorsDataBetweenTimePoints(httpClient IHttpAbstracter, startTimeUnix int64, endTimeUnix int64,
	timeofAverage string, sensorCodes []string) (viaCode SensorDataKeyedViaCode, viaTime SensorDataKeyedViaTime, err error) {

	bytesRead, err := httpClient.DoHttpGetCall(buildCompleteDataRequestURI(
		strconv.FormatInt(time.Now().Unix(), 10),
		strconv.FormatInt(startTimeUnix, 10),
		strconv.FormatInt(endTimeUnix, 10),
		timeofAverage, sensorCodes))
	if err != nil {
		return
	}

	pmProResponse, err := deserializePmProDataResponse(bytesRead)
	if err != nil {
		return
	}
	/*type pmProSensorsDataInTimePeriodResponse struct {
		EndTime   int64 `json:"end"`
		StartTime int64 `json:"start"`
		Data      [][]struct {
			Time  int64   `json:"t"` // "t": 1597611600, "v": 62.780000000000000,
			Value float32 `json:"v"`
		} `json:"values"`
		Vars []string `json:"vars"`
	}*/
	// "vars": [
	// 	"06HUMID_O:A1h",
	// 	"06PRESS_O:A1h",
	// 	"06PM10A_6_k:A1h",
	// 	"06PM25A_6_k:A1h"
	// 	]

	viaCode = processResponseViaSensorCode(pmProResponse)
	viaTime = SensorDataKeyedViaTime{Data: nil} //todo
	return
}

func processResponseViaSensorCode(r PmProSensorsDataInTimePeriodResponse) (result SensorDataKeyedViaCode) {
	result = SensorDataKeyedViaCode{
		Data: make([]sensorDataKeyedViaCode, len(r.Sensors)),
	}

	for sensorNameiterator, sensorData := range r.Data {
		singleSensorResult := sensorDataKeyedViaCode{
			SensorCode: r.Sensors[sensorNameiterator],
		}
		for _, singleDataEntry := range sensorData {
			singleSensorResult.Data = append(singleSensorResult.Data, struct {
				Time  int64   "json:\"time\""
				Value float32 "json:\"value\""
			}{Time: singleDataEntry.Time, Value: singleDataEntry.Value})
		}
		result.Data[sensorNameiterator] = singleSensorResult
	}
	return
}

/*func processResponseViaTime(r PmProSensorsDataInTimePeriodResponse) (result SensorDataKeyedViaTime) {
	result = SensorDataKeyedViaTime{
		Data: make([]sensorDataKeyedViaTime, len(r.Sensors)),
	}

	for sensorNameiterator, sensorData := range r.Data {
		sensorName := r.Sensors[sensorNameiterator]

		for _, singleDataEntry := range sensorData {
			singleSensorResult := sensorDataKeyedViaTime{
				Time: singleDataEntry.Time,
				Data: append(singleSensorResult.Data, struct {
					SensorCode string   "json:\"sensorCode\""
					Value float32 "json:\"value\""
				}{SensorCode: sensorName, Value: singleDataEntry.Value})
			}
			// singleSensorResult.Data = append(singleSensorResult.Data, struct {
			// 	SensorCode string   "json:\"time\""
			// 	Value float32 "json:\"value\""
			// }{SensorCode: singleDataEntry.Time, Value: singleDataEntry.Value})
		}
		result.Data[sensorNameiterator] = singleSensorResult
	}
	return
}*/
