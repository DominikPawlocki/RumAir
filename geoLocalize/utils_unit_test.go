package geolocalize

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// flag set up in Azure Build pipeline.
var withIntegration = flag.Bool("withIntegrationTests", true, "withIntegrationTests")

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

var localizedAirStationsMock = map[string]*LocalizedAirStation{
	"02":  &LocalizedAirStation{Station: localizableStation, CitiesNearby: []string{"aaa", "bbb"}},
	"021": &LocalizedAirStation{Station: notLocalizableStation, CitiesNearby: []string{"aaa", "ccc"}}}

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
