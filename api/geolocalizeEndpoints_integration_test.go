package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetStationNumbersPerCityHandler_ThenResponceIsCorrect(t *testing.T) {
	req, err := http.NewRequest("GET", "/stations/locate/locationIQ/numbersPerCity", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetStationNumbersPerCityHandler)
	handler.ServeHTTP(rr, req)

	var strigifiedResponse = rr.Body.String()
	assert.Equal(t, rr.Code, http.StatusOK, fmt.Sprintf("Handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK))

	//expectedResult := `["Gdansk with 6 stations : 27, 24, 4, 5, 21, 20","Gdynia with 16 stations : 25, 7, 8, 27, 24, 4, 26, 5, 21, 36, 6, 23, 29, 20, 35, 32","Plock with 3 stations : 37, 39, 38"]`
	assert.True(t, strings.Contains(strigifiedResponse, "Gdansk with "), fmt.Sprintln("Should contain \"Gdansk with \" string."))
	assert.True(t, strings.Contains(strigifiedResponse, "Gdynia with "), fmt.Sprintln("Should contain \"Gdynia with \" string."))
	assert.True(t, strings.Contains(strigifiedResponse, "Plock with "), fmt.Sprintln("Should contain \"Plock with \" string."))
}
