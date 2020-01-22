package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dompaw/RumAir/airStations"
	"github.com/stretchr/testify/assert"
)

func Test_ShowAllStationsSensorsCodesHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/stations/sensors/codes", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(MockableHTTPHandler{
		mockableDataFetcher: airStations.StationsCapabiltiesFetcher{},
		methodToBeCalled:    ShowAllStationsSensorsCodesHandler})
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK, fmt.Sprintf("Handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK))

	expected197chars :=
		`["Station : 00 can :  HES00_RH HES00_PA HES00_TA HES00_PM10 HES00_PM25","Station : 001 can :  001FLOW 001AR_AH 001AR_RH 001AR_TIN 001AR_TDEV 001NO 001NO2 001NOX 001RH 001PA 001RAIN 001TEMP 001WD`
	stringifiedResponse := rr.Body.String()

	assert.True(t, strings.HasPrefix(stringifiedResponse, expected197chars), fmt.Sprintf("expected first 197 signs like above, but was %v ", stringifiedResponse))
}

func Test_GetAllStationsCapabilitiesHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/stations/sensors", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(MockableHTTPHandler{
		mockableDataFetcher: airStations.StationsCapabiltiesFetcher{},
		methodToBeCalled:    GetAllStationsCapabilitiesHandler})
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK, fmt.Sprintf("Handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK))

	expectedBeggining :=
		`{"00":{"ID":0,"LatitudeSensor":"","LongitudeSensor":"","Sensors":[{"id":1759,"code":"HES00_RH","name":"RH HES061","compound_type":"humid","unit_id":"%","decimals":0,"format":"","average_type":"arithmetic","averages":"A10m,A30m,A1h","high_averages":"A24h,A1M,A1Y"},`
	assert.True(t, strings.HasPrefix(rr.Body.String(), expectedBeggining), fmt.Sprintf("expected first 197 signs like above, but was %v ", rr.Body.String()))
}
