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

func Test_Given_ErrorResponseFromDoAllMeasurmentsAPIcalll_When_GetAllStationsCapabilities_Then_Returns500WithErrorMessage(t *testing.T) {
	exampleMockErrorText := "timeout expired"

	sut := func() (*http.Request, error) {
		return http.NewRequest("GET", "/stations/sensors", nil)
	}

	rr := setUpMockAndPerformSutCall(t, nil, errors.New(exampleMockErrorText), sut)

	assert.Equal(t, http.StatusInternalServerError, rr.Code, fmt.Sprintf("Got %v code, want %v", rr.Code, http.StatusInternalServerError))

	bodyBytes, _ := ioutil.ReadAll(rr.Body)
	bodyString := string(bodyBytes)
	assert.True(t, strings.HasPrefix(bodyString, stationsCapabilitesFetchingError),
		fmt.Sprintf("Expected error starts like %s,but got %s in result", stationsCapabilitesFetchingError, bodyString))
	assert.Contains(t, bodyString, exampleMockErrorText)
}

func Test_Given_BrokenResponseFromDoAllMeasurmentsAPIcall_When_GetAllStationsCapabilities_Then_Returns500WithErrorMessage(t *testing.T) {
	exampleMockUnableToDeserializeResponse := []byte("bla bla bla ....")

	sut := func() (*http.Request, error) {
		return http.NewRequest("GET", "/stations/sensors", nil)
	}

	rr := setUpMockAndPerformSutCall(t, exampleMockUnableToDeserializeResponse, nil, sut)

	assert.Equal(t, http.StatusInternalServerError, rr.Code, fmt.Sprintf("Got %v code, want %v", rr.Code, http.StatusInternalServerError))

	bodyBytes, _ := ioutil.ReadAll(rr.Body)
	bodyString := string(bodyBytes)
	assert.True(t, strings.HasPrefix(bodyString, stationsCapabilitesFetchingError),
		fmt.Sprintf("Expected error starts like %s,but got %s in result", stationsCapabilitesFetchingError, bodyString))
}

func setUpMockAndPerformSutCall(t *testing.T, mockedResponse []byte, mockedError error, sut func() (*http.Request, error)) (rr *httptest.ResponseRecorder) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockIStationsCapabiltiesFetcher(ctrl)

	// Mock setting up
	m.
		EXPECT().
		DoAllMeasurmentsAPIcall().
		Return(mockedResponse, mockedError).
		AnyTimes()

	//req, _ := http.NewRequest("GET", "/stations/sensors", nil)

	req, _ := sut()

	rr = httptest.NewRecorder()
	handler := http.Handler(MockableHTTPHandler{mockableDataFetcher: m})
	handler.ServeHTTP(rr, req)
	return
}
