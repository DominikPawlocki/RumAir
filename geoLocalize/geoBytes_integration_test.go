package geolocalize

import (
	"fmt"
	"testing"

	"github.com/dompaw/RumAir_Pmpro_Sensors_API/airStations"
	"github.com/stretchr/testify/assert"
)

var localizableStation = &airStations.AirStation{
	ID:              2,
	LatitudeSensor:  "05LAT",
	LongitudeSensor: "05LON",
	Sensors: []airStations.SensorSimplified{
		airStations.SensorSimplified{
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
		airStations.SensorSimplified{
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

var notLocalizableStation = &airStations.AirStation{
	ID:              21,
	LatitudeSensor:  "",
	LongitudeSensor: "",
	Sensors: []airStations.SensorSimplified{
		airStations.SensorSimplified{
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

var expected = map[string]*airStations.AirStation{"02": localizableStation, "021": notLocalizableStation}

// 22/01/2020 - GeoBytes getNearbyCities seems not working !

func Test_LocalizeStationsGeoBytes(t *testing.T) {
	if !*withIntegration {
		t.Skip("Test ommited. Flag `withIntegrationTests` set to : false")
	}
	actual, _ := LocalizeStationsGeoBytes(expected)

	assert.Len(t, actual, 1, fmt.Sprintf("Station with id %v should be only localized, cause it has lat and lon 'sensors'.", localizableStation.ID))
	assert.Len(t, actual["02"].CitiesNearby, 2, "Wrong number of cities nearby")
	assert.ElementsMatchf(t, actual["02"].CitiesNearby, [2]string{"Gdynia", "Gdansk"}, "Wrong number of cities nearby.")
}
