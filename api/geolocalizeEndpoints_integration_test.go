package api

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// flag introduced, for possible distinguid differentciation from integration or unit tests. Not used right now.
//usage like : go test -v .\airStations\stationSensors.go .\airStations\utils.go .\airStations\stationSensors_integration_test.go  -args -isIntegration=true
var withIntegrationTests = flag.Bool("isIntegration", false, "isIntegration")

func TestMain(m *testing.M) {
	flag.Parse()
	fmt.Println("Flag `withIntegrationTests` set to : ", *withIntegrationTests)
}

func Test_GetStationNumbersPerCityHandler_ThenResponceIsCorrect(t *testing.T) {
	if !*withIntegrationTests {
		t.Skip("Takes about 1 min to pass..")
	}
	req, err := http.NewRequest("GET", "/stations/locate/locationIQ/numbersPerCity", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetStationNumbersPerCityHandler)
	handler.ServeHTTP(rr, req)

	var strigifiedResponse = rr.Body.String()
	assert.Equal(t, rr.Code, http.StatusOK, fmt.Sprintf("Handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK))

	assert.True(t, strings.Contains(strigifiedResponse, "Gdynia"), fmt.Sprintln("Should contain \"Gdansk with \" string."))
	assert.True(t, strings.Contains(strigifiedResponse, "Rumia"), fmt.Sprintln("Should contain \"Gdynia with \" string."))
	assert.True(t, strings.Contains(strigifiedResponse, "Płock"), fmt.Sprintln("Should contain \"Plock with \" string."))
}
