package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dompaw/RumAir_Pmpro_Sensors_API/airStations"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Test_GivenStationNumber04_WhenShowStationSensorCodesHandler_ThenResponceIsCorrect(t *testing.T) {
	if !*withIntegration {
		t.Skip("Test ommited. Flag `withIntegrationTests` set to : false")
	}
	req, err := http.NewRequest("GET", "/stations/04/sensors", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	//seriving it with router beacuse there is id parameter to be fetched
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.Handle("/stations/{id}/sensors", MockableHTTPHandler{
		mockableDataFetcher: airStations.HttpAbstracter{},
		methodToBeCalled:    GetStationSensorsHandler}).Methods("GET")
	myRouter.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK, fmt.Sprintf("Handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK))

	expected107chars :=
		`[{"id":87,"code":"04HUMID_F","name":"Wilgotność w komorze","compound_type":"humid","physical_device_id":19,`
	stringifiedResponse := rr.Body.String()

	assert.True(t, strings.HasPrefix(stringifiedResponse, expected107chars), fmt.Sprintf("expected first 197 signs like above, but was %v ", stringifiedResponse))
}

func Test_GivenIncorrectStationNumber_WhenShowStationSensorCodesHandlerHandler_ThenResponceIs404(t *testing.T) {
	if !*withIntegration {
		t.Skip("Test ommited. Flag `withIntegrationTests` set to : false")
	}
	req, err := http.NewRequest("GET", "/stations/incorrect/sensors", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	//seriving it with router beacuse there is id parameter to be fetched
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.Handle("/stations/{id}/sensors", MockableHTTPHandler{
		mockableDataFetcher: airStations.HttpAbstracter{},
		methodToBeCalled:    GetStationSensorsHandler}).Methods("GET")
	myRouter.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusNotFound, fmt.Sprintf("Handler returned wrong status code: got %v want %v", rr.Code, http.StatusNotFound))
}

func Test_WhenShowAllStationsSensorsCodesHandler_ThenResponceIsCorrect(t *testing.T) {
	if !*withIntegration {
		t.Skip("Test ommited. Flag `withIntegrationTests` set to : false")
	}
	req, err := http.NewRequest("GET", "/stations/sensors/codes", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(MockableHTTPHandler{
		mockableDataFetcher: airStations.HttpAbstracter{},
		methodToBeCalled:    ShowAllStationsSensorCodesHandler})
	handler.ServeHTTP(rr, req)

	stringifiedResponse := rr.Body.String()

	assert.Equal(t, rr.Code, http.StatusOK, fmt.Sprintf("Handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK))

	expected197chars :=
		`["Station : 00 can :  HES00_RH HES00_PA HES00_TA HES00_PM10 HES00_PM25","Station : 001 can :  001FLOW 001AR_AH 001AR_RH 001AR_TIN 001AR_TDEV 001NO 001NO2 001NOX 001RH 001PA 001RAIN 001TEMP 001WD`

	assert.True(t, strings.HasPrefix(stringifiedResponse, expected197chars), fmt.Sprintf("expected first 197 signs like above, but was %v ", stringifiedResponse))
}

func Test_WhenGetAllStationsCapabilitiesHandler_ThenResponceIsCorrect(t *testing.T) {
	if !*withIntegration {
		t.Skip("Test ommited. Flag `withIntegrationTests` set to : false")
	}
	req, err := http.NewRequest("GET", "/stations/sensors", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(MockableHTTPHandler{
		mockableDataFetcher: airStations.HttpAbstracter{},
		methodToBeCalled:    GetAllStationsCapabilitiesHandler})
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK, fmt.Sprintf("Handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK))

	expectedBeggining :=
		`{"00":{"ID":0,"LatitudeSensor":"","LongitudeSensor":"","Sensors":[{"id":1759,"code":"HES00_RH","name":"RH HES061","compound_type":"humid","unit_id":"%","decimals":0,"format":"","average_type":"arithmetic","averages":"A10m,A30m,A1h","high_averages":"A24h,A1M,A1Y"},`
	assert.True(t, strings.HasPrefix(rr.Body.String(), expectedBeggining), fmt.Sprintf("expected first 197 signs like above, but was %v ", rr.Body.String()))
}
