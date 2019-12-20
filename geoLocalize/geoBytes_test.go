package geolocalize

import (
	"testing"

	"github.com/dompaw/RumAir/airStations"
	"github.com/stretchr/testify/assert"
)

var expectedAirStation1 = &airStations.AirStation{
	ID:              2,
	LatitudeSensor:  "04LAT",
	LongitudeSensor: "04LON",
	//Sensors:         make([]airStations.SensorMeasurmentSimpleType, 2),
	Sensors: []airStations.SensorMeasurmentSimpleType{
		airStations.SensorMeasurmentSimpleType{
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
		airStations.SensorMeasurmentSimpleType{
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

var expectedAirStation2 = &airStations.AirStation{
	ID:              21,
	LatitudeSensor:  "",
	LongitudeSensor: "",
	Sensors: []airStations.SensorMeasurmentSimpleType{
		airStations.SensorMeasurmentSimpleType{
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

var expected = map[string]*airStations.AirStation{"02": expectedAirStation1, "021": expectedAirStation2}

func Test_LocalizeStationsGeoBytes(t *testing.T) {
	actual, _ := LocalizeStationsGeoBytes(expected)

	assert.Len(t, actual, 1, "Wrong lenght")
	assert.Len(t, actual["02"].CitiesNearby, 2, "Wrong number of books.")
	assert.ElementsMatchf(t, actual["02"].CitiesNearby, [2]string{"Gdynia", "Gdansk"}, "Wrong number of books.")
}
