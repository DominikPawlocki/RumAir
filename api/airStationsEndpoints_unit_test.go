package api

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_Given_ErrorResponseFromPmProApiCall_When_GetAllStationsCapabilities_Then_Returns500WithErrorMessage(t *testing.T) {
	exampleMockErrorText := "timeout expired"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockIStationsCapabiltiesFetcher(ctrl)

	// Mock setting up
	m.
		EXPECT().
		DoAllMeasurmentsAPIcall().
		Return(nil, errors.New(exampleMockErrorText)).
		AnyTimes()

	req, _ := http.NewRequest("GET", "/stations/sensors", nil)

	rr := httptest.NewRecorder()
	handler := http.Handler(MockableHTTPHandler{mockableDataFetcher: m})
	handler.ServeHTTP(rr, req)

	assert.NotNil(t, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code, fmt.Sprintf("Got %v code, want %v", rr.Code, http.StatusInternalServerError))

	bodyBytes, _ := ioutil.ReadAll(rr.Body)
	bodyString := string(bodyBytes)
	assert.True(t, strings.HasPrefix(bodyString, stationsCapabilitesFetchingError),
		fmt.Sprintf("Expected error starts like %s,but got %s in result", stationsCapabilitesFetchingError, bodyString))
}

func Test_Given_BrokenResponseFromPmProApiCall_When_GetAllStationsCapabilities_Then_Returns500WithErrorMessage(t *testing.T) {
	exampleMockUnableToDeserializeResponse := []byte("bla bla bla ....")

	rr := setUpMockAndPerformSutCall(t, exampleMockUnableToDeserializeResponse, nil)

	assert.Equal(t, http.StatusInternalServerError, rr.Code, fmt.Sprintf("Got %v code, want %v", rr.Code, http.StatusInternalServerError))

	bodyBytes, _ := ioutil.ReadAll(rr.Body)
	bodyString := string(bodyBytes)
	assert.True(t, strings.HasPrefix(bodyString, stationsCapabilitesFetchingError),
		fmt.Sprintf("Expected error starts like %s,but got %s in result", stationsCapabilitesFetchingError, bodyString))
}

func setUpMockAndPerformSutCall(t *testing.T, mockedResponse ...interface{}) (rr *httptest.ResponseRecorder) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockIStationsCapabiltiesFetcher(ctrl)

	// Mock setting up
	m.
		EXPECT().
		DoAllMeasurmentsAPIcall().
		Return(mockedResponse). //zwroc 1wszy i drugi
		AnyTimes()

	req, _ := http.NewRequest("GET", "/stations/sensors", nil)

	rr = httptest.NewRecorder()
	handler := http.Handler(MockableHTTPHandler{mockableDataFetcher: m})
	handler.ServeHTTP(rr, req)
	return
}
