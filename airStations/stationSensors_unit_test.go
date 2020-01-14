package airStations

import (
	"errors"
	"fmt"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_Given_APIError_When_GetAllStationsCapabilities_Then_ResponseIsNilDataAndError(t *testing.T) {
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

	stations, err = GetAllStationsCapabilities(m)

	assert.Nil(t, stations)
	assert.Equal(t, errorText, err.Error(), fmt.Sprintf("Expected error like %s,but got %s in result", errorText, err.Error()))
}
