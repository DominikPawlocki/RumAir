package airStations

import (
	"errors"
	"fmt"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_Given_ErrorResponseFromPmProApiCall_When_GetAllStationsCapabilities_Then_ResponseIsNilAndError(t *testing.T) {
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

	stations, err = GetAllStationsCapabilities(m)

	assert.Nil(t, stations)
	assert.Equal(t, exampleMockErrorText, err.Error(), fmt.Sprintf("Expected error like %s,but got %s in result", exampleMockErrorText, err.Error()))
}
