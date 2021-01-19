package airStations

import (
	"strconv"
	"time"
)

type SensorDataKeyedViaTime struct {
	Data []sensorDataKeyedViaTime `json:"data"`
}

type sensorDataKeyedViaTime struct {
	TimeUnix int64     `json:"tUnix"`
	Time     time.Time `json:"t"`
	Data     []struct {
		SensorCode string  `json:"c"`
		Value      float32 `json:"v"`
	} `json:"data"`
}

type SensorDataKeyedViaCode struct {
	Data []sensorDataKeyedViaCode `json:"data"`
}

type sensorDataKeyedViaCode struct {
	SensorCode string `json:"sensorCode"`
	Data       []struct {
		TimeUnix int64     `json:"tUnix"`
		Time     time.Time `json:"t"`
		Value    float32   `json:"v"`
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
	viaTime = processResponseViaTime(pmProResponse)
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
				TimeUnix int64     "json:\"tUnix\""
				Time     time.Time "json:\"t\""
				Value    float32   "json:\"v\""
			}{TimeUnix: singleDataEntry.Time, Time: time.Unix(singleDataEntry.Time, 0), Value: singleDataEntry.Value})
		}
		result.Data[sensorNameiterator] = singleSensorResult
	}
	return
}

func processResponseViaTime(r PmProSensorsDataInTimePeriodResponse) (result SensorDataKeyedViaTime) {
	type readingInTimeTemp struct {
		SensorCode string
		Value      float32
	}

	var byTime = make(map[int64][]*readingInTimeTemp)

	for sensorNameiterator, sensorData := range r.Data {
		sensorCode := r.Sensors[sensorNameiterator]
		for _, singleDataEntry := range sensorData {
			byTime[singleDataEntry.Time] = append(byTime[singleDataEntry.Time], &readingInTimeTemp{SensorCode: sensorCode, Value: singleDataEntry.Value})
		}
	}

	result = SensorDataKeyedViaTime{}
	for timeUnix, readings := range byTime {
		r := sensorDataKeyedViaTime{TimeUnix: timeUnix, Time: time.Unix(timeUnix, 0)}
		for _, singleRead := range readings {
			r.Data = append(r.Data, struct {
				SensorCode string  `json:"c"`
				Value      float32 `json:"v"`
			}{SensorCode: singleRead.SensorCode, Value: singleRead.Value})
		}
		result.Data = append(result.Data, r)
	}
	return
}
