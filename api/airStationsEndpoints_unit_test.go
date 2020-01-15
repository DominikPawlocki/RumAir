package api

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_Whem_HttpErrorResponseFromPmPro_Then_GetAllStationsCapabilities_Returns500AndError(t *testing.T) {
	errorText := "timeout expired"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockIStationsCapabiltiesFetcher(ctrl)

	// Mock setting up
	m.
		EXPECT().
		DoAllMeasurmentsAPIcall().
		Return(nil, errors.New(errorText)).
		AnyTimes()

	req, _ := http.NewRequest("GET", "/stations/sensors", nil)

	rr := httptest.NewRecorder()
	handler := http.Handler(MockableHTTPHandler{mockableDataFetcher: m})
	handler.ServeHTTP(rr, req)

	assert.NotNil(t, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code, fmt.Sprintf("Handler returned wrong status code: got %v want %v", rr.Code, http.StatusInternalServerError))
	// zczytaj body requesta zamiast !!!

	assert.Equal(t, stationsCapabilitesFetchingError, rr.b.Error(), fmt.Sprintf("Expected error like %s,but got %s in result", stationsCapabilitesFetchingError, err.Error()))

}
