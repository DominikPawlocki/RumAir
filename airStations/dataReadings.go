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

type PmProErrorResponse struct{
	<!DOCTYPE HTML>
	<html>
	<head>
	<title>Server Error: Error during processing /data/averages handler can&#39;t read &amp;quot...</title>
	<style type='text/css'>
		  html * { padding:0; margin:0; }
		  ....
		</style>
	</head>
	<body>
	<h2>Server Error: Error during processing /data/averages handler can&#39;t read &quot;codeMap(08PRo...</h2>
	
	<div id='summary'>Error during processing /data/averages handler can&#39;t read &quot;codeMap(08PRoESS_O)&quot;: no such element in array
		while executing
	&quot;dict get &#x24;codeMap(&#x24;keyPart1) id&quot;
		(procedure &quot;_/data/averages&quot; line 35)
		invoked from within
	&quot;_/data/averages &#x24;req &#x7B;*&#x7D;&#x24;args &quot;</div>
			<div id='errorinfo'><p>Caller: <code>::Http::ServerError &#x7B;-host 127.0.0.1 -port 12000 -httpd &#x7B;::Httpd new&#x7D; -id ::oo::Obj248 -server &#x7B;127.0.0.1 127.0.0.1 12000&#x7D; -scheme http -sock sock7f5a90e50fd0 -cid ::oo::Obj24873 -ipaddr 127.0.0.1 -rport 54656 -received_seconds 1610837982 -server_id &#x7B;Wub 6.0&#x7D; -pipeline 12857 -send ::Httpd::coros::obj24873 -transaction 1 -time &#x7B;connected 1610837982962994&#x7D; -header &#x7B;GET /webapp/data/averages?_dc=1571382648880&amp;type=chart&amp;avg=1h&amp;start=1571123328&amp;end=1571382528&amp;vars=08HUMID_O%3AA1h%2C08PRoESS_O%3AA1h%2C08PM10A_6_k%3AA1h%2C08PM25A_6_k%3AA1h HTTP/1.0&#x7D; -method GET host 127.0.0.1:12000 -clientheaders &#x7B;host connection user-agent accept accept-language accept-encoding upgrade-insecure-requests cache-control&#x7D; user-agent &#x7B;Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:84.0) Gecko/20100101 Firefox/84.0&#x7D; accept &#x7B;text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8&#x7D; accept-language &#x7B;en-GB,en;q=0.5&#x7D; accept-encoding &#x7B;gzip, deflate, br&#x7D; upgrade-insecure-requests 1 -version 1.0 -ua &#x7B;ua &#x7B;Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:84.0) Gecko/20100101 Firefox/84.0&#x7D; id FF version &#x7B;&#x7D; mozilla_version 5.0 extensions &#x7B;&#x7D; platform &#x7B;Windows NT 10.0&#x7D; security Win64 subplatform x64 language rv:84.0 product &#x7B;Gecko 20100101 Firefox 84.0&#x7D;&#x7D; -ua_class unknown -normalized 1 -path /webapp/data/averages -query _dc=1571382648880&amp;type=chart&amp;avg=1h&amp;start=1571123328&amp;end=1571382528&amp;vars=08HUMID_O%3AA1h%2C08PRoESS_O%3AA1h%2C08PM10A_6_k%3AA1h%2C08PM25A_6_k%3AA1h -url http://127.0.0.1:12000/webapp/data/averages -uri http://127.0.0.1:12000/webapp/data/averages?_dc=1571382648880&amp;type=chart&amp;avg=1h&amp;start=1571123328&amp;end=1571382528&amp;vars=08HUMID_O%3AA1h%2C08PRoESS_O%3AA1h%2C08PM10A_6_k%3AA1h%2C08PM25A_6_k%3AA1h -cache-control max-age=0 -forwards &#x7B;&#x7D; -encoding binary -received 1610837982963956 -cookies &#x7B;&#x7D; -Query &#x7B;_dc &#x7B;1571382648880 &#x7B;-count 1&#x7D;&#x7D; type &#x7B;chart &#x7B;-count 2&#x7D;&#x7D; avg &#x7B;1h &#x7B;-count 3&#x7D;&#x7D; start &#x7B;1571123328 &#x7B;-count 4&#x7D;&#x7D; end &#x7B;1571382528 &#x7B;-count 5&#x7D;&#x7D; vars &#x7B;08HUMID_O:A1h,08PRoESS_O:A1h,08PM10A_6_k:A1h,08PM25A_6_k:A1h &#x7B;-count 6&#x7D;&#x7D;&#x7D; -prefix /webapp -suffix data/averages -extension &#x7B;&#x7D; content-type x-text/html-fragment -dynamic 1 -extra &#x7B;&#x7D; -fprefix data/averages -cprefix data/averages&#x7D; &#x7B;Error during processing /data/averages handler can&#39;t read &quot;codeMap(08PRoESS_O)&quot;: no such element in array
		while executing

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
