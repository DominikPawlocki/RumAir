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
		ID:              2,
		LatitudeSensor:  "04LAT",
		LongitudeSensor: "04LON",
		Sensors: []SensorMeasurmentSimpleType{
			SensorMeasurmentSimpleType{
				ID:           1760,
				Code:         "HES00_PA",
				Name:         "PA HES061",
				CompoundType: "press",
				UnitID:       "hPa",
				Decimals:     0,
				Format:       "",
				AverageType:  "arithmetic",
				Averages:     "A10m,A30m,A1h",
				HighAverages: "A24h,A1M,A1Y"},
			SensorMeasurmentSimpleType{
				ID:           1760,
				Code:         "HES00_PA",
				Name:         "PA HES061",
				CompoundType: "press",
				UnitID:       "hPa",
				Decimals:     0,
				Format:       "",
				AverageType:  "arithmetic",
				Averages:     "A10m,A30m,A1h",
				HighAverages: "A24h,A1M,A1Y"}},
		SensorsCount: 2}

	var notLocalizableStation = &AirStation{
		ID:              21,
		LatitudeSensor:  "",
		LongitudeSensor: "",
		Sensors: []SensorMeasurmentSimpleType{
			SensorMeasurmentSimpleType{
				ID:           1760,
				Code:         "HES00_PA",
				Name:         "PA HES061",
				CompoundType: "press",
				UnitID:       "hPa",
				Decimals:     0,
				Format:       "",
				AverageType:  "arithmetic",
				Averages:     "A10m,A30m,A1h",
				HighAverages: "A24h,A1M,A1Y"}},
		SensorsCount: 1}

	var stations = map[string]*AirStation{"02": localizableStation, "021": notLocalizableStation}

	actual := ShowStationsSensorsCodes(stations)

	
	[0]:"Station : 02 can :  HES00_PA HES00_PA"
	[1]:"Station : 021 can :  HES00_PA"

	assert.GreaterOrEqual(t, len(actual), 30, fmt.Sprintf("There should be minimum like 30 stations fetched. %v stations was fetched.", len(actual)))
	assert.GreaterOrEqual(t, len(actual[5]), 70, fmt.Sprintf("Station with slice index 5 should have more sensors. Now it has %v", actual[5]))

	assert.Condition(t, func() (success bool) { return sort.StringsAreSorted(actual) }, fmt.Sprint("Should be sorted."))
}
