package airStations

import (
	"errors"
	"fmt"
	"sort"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_Given_ErrorResponseFromDoAllMeasurmentsAPIcall_When_GetAllStationsCapabilities_Then_ResponseIsNilAndError(t *testing.T) {
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

func Test_GivenTwoStations_When_ShowStationsSensorsCodes_Then_(t *testing.T) {
	var localizableStation = &AirStation{
		ID: 2,
		Sensors: []SensorMeasurmentSimpleType{
			SensorMeasurmentSimpleType{
				ID:   1760,
				Code: "Sensor02_AA"},
			SensorMeasurmentSimpleType{
				ID:   1760,
				Code: "Sensor02_BB"}},
		SensorsCount: 2}

	var notLocalizableStation = &AirStation{
		ID: 21,
		Sensors: []SensorMeasurmentSimpleType{
			SensorMeasurmentSimpleType{
				ID:   1760,
				Code: "HES21_PA"}},
		SensorsCount: 1}

	var stations = map[string]*AirStation{"02": localizableStation, "021": notLocalizableStation}

	actual := ShowStationsSensorsCodes(stations)

	var expectedFirst = "Station : 02 can :  Sensor02_AA Sensor02_BB"
	var expectedSecond = "Station : 021 can :  HES21_PA"

	assert.Equal(t, expectedFirst, actual[0], fmt.Sprintf("There should %s and is %s", expectedFirst, actual[0]))
	assert.Equal(t, expectedSecond, actual[1], fmt.Sprintf("There should %s and is %s", expectedSecond, actual[1]))

	assert.Condition(t, func() (success bool) { return sort.StringsAreSorted(actual) }, fmt.Sprint("Should be sorted."))
}

func Test_Given_StationNumber04_And_ErrorResponseFromDoAllMeasurmentsAPIcall_When_GetStationSensors_Then_ResponseIsNilAndError(t *testing.T) {
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

	actual := GetStationSensors(StationsCapabiltiesFetcher{}, "04")

	assert.Nil(t, stations)
	assert.Equal(t, exampleMockErrorText, err.Error(), fmt.Sprintf("Expected error like %s,but got %s in result", exampleMockErrorText, err.Error()))
}
