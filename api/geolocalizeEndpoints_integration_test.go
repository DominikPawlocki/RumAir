package api

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// flag set up in Azure Build pipeline.
var withIntegration = flag.Bool("withIntegrationTests", true, "withIntegrationTests")

func TestMain(m *testing.M) {
	flag.Parse()
	var msg string = "RUNNING"

	if !*withIntegration {
		msg = "OMMITING"
	}

	fmt.Printf("%s integration tests in module `api`. Flag `withIntegrationTests` set to : %v \n", msg, *withIntegration)
	code := m.Run()
	os.Exit(code)
}

func Test_GetStationNumbersPerCityHandler_ThenResponceIsCorrect(t *testing.T) {
	if !*withIntegration {
		t.Skip("Test ommited. Flag `withIntegrationTests` set to : false")
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
	fmt.Printf(strigifiedResponse)

	assert.True(t, strings.Contains(strigifiedResponse, "Gdynia"), fmt.Sprintln("Should contain \"Gdynia\" string."))
	assert.True(t, strings.Contains(strigifiedResponse, "Rumia"), fmt.Sprintln("Should contain \"Rumia\" string."))
	assert.True(t, strings.Contains(strigifiedResponse, "Płock"), fmt.Sprintln("Should contain \"Plock\" string."))
}
