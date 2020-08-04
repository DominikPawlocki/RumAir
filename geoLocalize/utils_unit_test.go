package geolocalize

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/dompaw/RumAir_Pmpro_Sensors_API/airStations"
	"github.com/stretchr/testify/assert"
)

// flag set up in Azure Build pipeline.
var withIntegration = flag.Bool("withIntegrationTests", false, "withIntegrationTests")

func TestMain(m *testing.M) {
	flag.Parse()
	var msg string = "RUNNING"

	if !*withIntegration {
		msg = "OMMITING"
	}

	fmt.Printf("%s integration tests in module `geoLocalize`. Flag `withIntegrationTests` set to : %v \n", msg, *withIntegration)
	code := m.Run()
	os.Exit(code)
}

var localizableStationSimplified = &airStations.AirStationSimplified{
	ID:           2,
	SensorsCount: 2}

var notLocalizableStationSimplified = &airStations.AirStationSimplified{
	ID:           21,
	SensorsCount: 1}

var localizedAirStationsMock = map[string]*LocalizedAirStationSimplified{
	"02":  &LocalizedAirStationSimplified{Station: localizableStationSimplified, CitiesNearby: []string{"aaa", "bbb"}},
	"021": &LocalizedAirStationSimplified{Station: notLocalizableStationSimplified, CitiesNearby: []string{"aaa", "ccc"}}}

func Test_GetStationNrPerCity(t *testing.T) {
	result := GetStationNrPerCity(localizedAirStationsMock)

	assert.Len(t, result, 3)
	assert.Equal(t, "aaa", result[0].City)
	assert.Contains(t, result[0].StationIdsConcat, "2")
	assert.Contains(t, result[0].StationIdsConcat, "21")

	assert.Equal(t, "bbb", result[1].City)
	assert.Equal(t, "2", result[1].StationIdsConcat)

	assert.Equal(t, "ccc", result[2].City)
	assert.Equal(t, "21", result[2].StationIdsConcat)
}
