package airStations

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"io/ioutil"
	"net/http"
)

var allStationsMeasurmentsURL string = "http://pmpro.dacsystem.pl/webapp/json/do?table=Measurement&v=2"
var pmproSystemBaseAPIURL string = "http://pmpro.dacsystem.pl/webapp/data"

type IHttpAbstracter interface {
	DoHttpGetCall(uri string) ([]byte, error)
}
type HttpAbstracter struct {
}

func (HttpAbstracter) DoHttpGetCall(uri string) ([]byte, error) {
	return doHttpGetCall(uri)
}

func doHttpGetCall(uri string) (bytesRead []byte, err error) {
	var netResp *http.Response

	netResp, err = http.Get(uri)
	if err != nil {
		return
	}

	defer netResp.Body.Close()

	//SLICE INITIALIZATIONS !
	//allMeasurments := make([]Sensor, 2)
	//var allMeasurments *[]Sensor = &[]Sensor{}

	bytesRead, err = ioutil.ReadAll(netResp.Body)
	if err != nil {
		fmt.Printf("Error during ReadAll bytesRead: %s err: %v. \n", bytesRead, err)
		return
	}

	fmt.Printf("%v bytes read from network for %s endpoint. \n", len(bytesRead), uri)
	return
}

func DoHttpCallWithConsoleDots(fn func(uri string) ([]byte, error), uri string) (bytesRead []byte, err error) {
	ticker := time.NewTicker(100 * time.Millisecond)
	done := make(chan bool)

	go selectTicker(done, ticker)

	bytesRead, err = fn(uri)

	ticker.Stop()
	done <- true
	return
}

/// Averages or high averages. Please be avare that particular sensors can have particular averages available
const (
	delimiter3A = "%3A" //:
	delimiter2C = "%2C"
	A10M        = "A10m"
	A30M        = "A30m"
	A1H         = "A1h"
	A24H        = "A24h"
	A8H         = "A8h"
	A8hMax      = "A8h_max"
	SA8H        = "sA8h"
	SA8hMax     = "sA8h_max"
	//,A8h_max
	//,A1M,A1Y,X24h_8h,X1M_Q24h",
)

/* the task is to create URI like :
https://pmpro.dacsystem.pl/webapp/data/averages?_dc=1571382648880&type=chart&avg=1h&start=1571123328&end=1571382528&vars=08HUMID_O%3AA1h%2C08PRESS_O%3AA1h%2C08PM10A_6_k%3AA1h%2C08PM25A_6_k%3AA1h
*/
func buildCompleteDataRequestURI(requestTime string, startTime string, endTime string, timeOfAverage string, sensorCodes []string) (uri string) {
	var strBldr strings.Builder

	strBldr.WriteString(pmproSystemBaseAPIURL)
	strBldr.WriteString("/averages")
	strBldr.WriteString("?_dc=")
	strBldr.WriteString(requestTime) //fmt.Sprintf("dasda %s, time.Now().Unix())
	strBldr.WriteString("&type=chart")
	//strBldr.WriteString("&avg=1h") // CHECK
	strBldr.WriteString("&start=")
	strBldr.WriteString(startTime)
	strBldr.WriteString("&end=")
	strBldr.WriteString(endTime)
	strBldr.WriteString("&vars=")
	strBldr.WriteString(buildDataRequestSensorsURIPart(timeOfAverage, sensorCodes))

	return strBldr.String()
}

//08HUMID_O%3AA1h%2C08PRESS_O%3AA1h%2C08PM10A_6_k%3AA1h%2C08PM25A_6_k%3AA1h
func buildDataRequestSensorsURIPart(timeOfAverage string, sensorCodes []string) string {
	var strBldr strings.Builder
	var sensorCodesNr = len(sensorCodes)
	for i, sensorCode := range sensorCodes {
		strBldr.WriteString(sensorCode)
		strBldr.WriteString(delimiter3A)
		strBldr.WriteString(timeOfAverage)
		if i < sensorCodesNr-1 {
			strBldr.WriteString(delimiter2C)
		}
	}

	return strBldr.String()
}

func deserializePmProDataResponse(bytesRead []byte) (pmProResponse PmProSensorsDataInTimePeriodResponse, err error) {
	err = json.Unmarshal(bytesRead, &pmProResponse)
	if err != nil {
		//System works that way, then when error occures, there is a view returned as 200 response.
		fmt.Println("Error during deserializing occured :", err)
		return
	}
	return pmProResponse, nil
}
