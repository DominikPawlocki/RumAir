package airStations

import (
	"strconv"
	"time"
)

type SensorDataKeyedViaTime struct {
	Time int64 `json:"time"`
	Data []struct {
		SensorCode string  `json:"sensorCode"`
		Value      float32 `json:"value"`
	} `json:"data"`
}

type SensorDataKeyedViaSensorCodeResponse struct {
	Data []sensorDataKeyedViaSensorCodeResponse `json:"data"`
}

type sensorDataKeyedViaSensorCodeResponse struct {
	SensorCode string `json:"sensorCode"`
	Data       []struct {
		Time  int64   `json:"time"`
		Value float32 `json:"value"`
	} `json:"data"`
}

func GetSensorsDataBetweenTimePointsVIATIME(httpClient IHttpAbstracter, startTime time.Time, endTime time.Time, timeofAverage string, sensorCodes []string) (result SensorDataKeyedViaSensorCodeResponse, err error) {
	result = SensorDataKeyedViaSensorCodeResponse{}
	result.Data = make([]sensorDataKeyedViaSensorCodeResponse, len(sensorCodes))

	bytesRead, err := httpClient.DoHttpGetCall(buildCompleteDataRequestURI(
		strconv.FormatInt(time.Now().Unix(), 10),
		strconv.FormatInt(startTime.Unix(), 10),
		strconv.FormatInt(endTime.Unix(), 10),
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
		Data      []struct {
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

	singleSensorResult := sensorDataKeyedViaSensorCodeResponse{}

	i := 0
	for sensorCode, sensorData := range pmProResponse.Data {
		singleSensorResult.SensorCode = pmProResponse.Vars[sensorCode]
		for i, oneData := range sensorData {
			singleSensorResult.Data[i].Time = oneData.Time
			singleSensorResult.Data[i].Value = oneData.Value

		}
		result.Data[i] = singleSensorResult
		i++
	}
	return
}
