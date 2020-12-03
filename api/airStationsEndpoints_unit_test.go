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

func Test_Given_ErrorResponseFromDoAllMeasurmentsAPICall_And_StationNumberCorrect_When_ShowStationSensorCodesHandler_Then_Returns404(t *testing.T) {
	exampleMockErrorText := "timeout expired"

	sut := func() (*http.Request, error) {
		return http.NewRequest("GET", "/stations/04/sensors", nil)
	}

	mock := setUpMock(t, nil, errors.New(exampleMockErrorText))

	req, _ := sut()

	rr := httptest.NewRecorder()
	handler := http.Handler(MockableHTTPHandler{mockableDataFetcher: mock, methodToBeCalled: GetStationSensorsHandler})
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code, fmt.Sprintf("Got %v code, want %v", rr.Code, http.StatusNotFound))
}

func Test_Given_ErrorResponseFromDoAllMeasurmentsAPIcalll_When_GetAllStationsCapabilities_Then_Returns500WithErrorMessage(t *testing.T) {
	exampleMockErrorText := "timeout expired"

	sut := func() (*http.Request, error) {
		return http.NewRequest("GET", "/stations/sensors", nil)
	}

	mock := setUpMock(t, nil, errors.New(exampleMockErrorText))

	req, _ := sut()

	rr := httptest.NewRecorder()
	handler := http.Handler(MockableHTTPHandler{mockableDataFetcher: mock, methodToBeCalled: GetAllStationsCapabilitiesHandler})
	handler.ServeHTTP(rr, req)

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

	mock := setUpMock(t, exampleMockUnableToDeserializeResponse, nil)

	req, _ := sut()

	rr := httptest.NewRecorder()
	handler := http.Handler(MockableHTTPHandler{mockableDataFetcher: mock, methodToBeCalled: GetAllStationsCapabilitiesHandler})
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code, fmt.Sprintf("Got %v code, want %v", rr.Code, http.StatusInternalServerError))

	bodyBytes, _ := ioutil.ReadAll(rr.Body)
	bodyString := string(bodyBytes)
	assert.True(t, strings.HasPrefix(bodyString, stationsCapabilitesFetchingError),
		fmt.Sprintf("Expected error starts like %s,but got %s in result", stationsCapabilitesFetchingError, bodyString))
}

func Test_Given_EmptyResponseFromDoAllMeasurmentsAPIcall_When_GetAllStationsCapabilitiesHandler_Then_Returns500WithErrorMessage(t *testing.T) {
	var emptyResponse []byte = make([]byte, 0)

	sut := func() (*http.Request, error) {
		return http.NewRequest("GET", "/stations/sensors/codes", nil)
	}

	mock := setUpMock(t, emptyResponse, nil)

	req, _ := sut()

	rr := httptest.NewRecorder()
	handler := http.Handler(MockableHTTPHandler{mockableDataFetcher: mock, methodToBeCalled: GetAllStationsCapabilitiesHandler})
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code, fmt.Sprintf("Got %v code, want %v", rr.Code, http.StatusInternalServerError))

	bodyBytes, _ := ioutil.ReadAll(rr.Body)
	bodyString := string(bodyBytes)
	assert.True(t, strings.HasPrefix(bodyString, stationsCapabilitesFetchingError),
		fmt.Sprintf("Expected error starts like %s,but got %s in result", stationsCapabilitesFetchingError, bodyString))
	assert.Contains(t, bodyString, "unexpected end of JSON input")
}

func Test_Given_ErrorResponseFromDoAllMeasurmentsAPIcall_When_ShowAllStationsSensorsCodesHandler_Then_Returns500WithErrorMessage(t *testing.T) {
	exampleMockErrorText := "timeout expired"

	sut := func() (*http.Request, error) {
		return http.NewRequest("GET", "/stations/sensors/codes", nil)
	}

	mock := setUpMock(t, nil, errors.New(exampleMockErrorText))

	req, _ := sut()

	rr := httptest.NewRecorder()
	handler := http.Handler(MockableHTTPHandler{mockableDataFetcher: mock, methodToBeCalled: ShowAllStationsSensorCodesHandler})
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code, fmt.Sprintf("Got %v code, want %v", rr.Code, http.StatusInternalServerError))

	bodyBytes, _ := ioutil.ReadAll(rr.Body)
	bodyString := string(bodyBytes)
	assert.True(t, strings.HasPrefix(bodyString, stationsCapabilitesFetchingError),
		fmt.Sprintf("Expected error starts like %s,but got %s in result", stationsCapabilitesFetchingError, bodyString))
	assert.Contains(t, bodyString, exampleMockErrorText)
}

func Test_Given_CorrectResponseFromDoAllMeasurmentsAPIcall_When_ShowAllStationsSensorsCodesHandler_Then_Returns200(t *testing.T) {
	//strings in `` instead of "" are auto-escaped !
	faultyResponse := []byte(`{"success":true,"totalCount":1655,"message":"","data": [	{123"694}	]}`)

	sut := func() (*http.Request, error) {
		return http.NewRequest("GET", "/stations/sensors/codes", nil)
	}

	mock := setUpMock(t, faultyResponse, nil)

	req, _ := sut()

	rr := httptest.NewRecorder()
	handler := http.Handler(MockableHTTPHandler{mockableDataFetcher: mock, methodToBeCalled: ShowAllStationsSensorCodesHandler})
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code, fmt.Sprintf("Got %v code, want %v", rr.Code, http.StatusInternalServerError))

	bodyBytes, _ := ioutil.ReadAll(rr.Body)
	bodyString := string(bodyBytes)
	assert.True(t, strings.HasPrefix(bodyString, stationsCapabilitesFetchingError),
		fmt.Sprintf("Expected error starts like %s,but got %s in result", stationsCapabilitesFetchingError, bodyString))
}

func setUpMock(t *testing.T, mockedResponse []byte, mockedError error) (m *MockIHttpAbstracter) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m = NewMockIHttpAbstracter(ctrl)

	// Mock setting up
	m.
		EXPECT().
		DoHttpGetCall("http://pmpro.dacsystem.pl/webapp/json/do?table=Measurement&v=2").
		Return(mockedResponse, mockedError).
		AnyTimes()

	return
}
