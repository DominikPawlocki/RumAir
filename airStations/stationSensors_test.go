package airStations

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var stations map[string]*AirStation
var err error

func setup() {
	stations, err = GetAllStationsCapabilities()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	//shutdown()
	os.Exit(code)
}

func Test_GetAllStationsCapabilities(t *testing.T) {
	assert.Nil(t, err, fmt.Sprintf("There is error %v: ", err))

	assert.GreaterOrEqual(t, len(stations), 30, fmt.Sprintf("There should be minimum like 30 stations fetched. %v stations was fetched. ", len(stations)))
	assert.Contains(t, stations, "02", "Station '02'should exist. ")
	assert.GreaterOrEqual(t, len(stations["02"].Sensors), 25, fmt.Sprintf("Station '02' had like 10 sensors minimum. Now it has %v", len(stations["02"].Sensors)))
}

func Test_GetAllStationsCapabilities_StationShouldContainOnlyOwnSensors(t *testing.T) {
	for idx, station := range stations {
		for _, sensor := range station.Sensors {
			assert.Truef(t, strings.HasPrefix(sensor.Code, idx),
				fmt.Sprintf("Every sensor for station should start with its id. %s in that case, and sensor starts wrongly with %s.", idx, sensor.Code))
		}
	}
}

func Test_GetStationsCapabilities_StationShouldContainOnlyOwnSensors(t *testing.T) {
	var stationID string = "04"
	actual := GetStationCapabilities(stationID)

	assert.NotNil(t, actual)
	assert.Equal(t, actual.ID, 4, fmt.Sprintf("Asking for %v stationID, got %v in result", stationID, actual.ID))
	for _, sensor := range actual.Sensors {
		assert.Truef(t, strings.HasPrefix(sensor.Code, stationID),
			fmt.Sprintf("Every sensor for station should start with its id. %s in that case, and sensor starts wrongly with %s.", stationID, sensor.Code))
	}
}

func Test_ShowStationsSensorsCodes(t *testing.T) {
	actual := ShowStationsSensorsCodes(stations)

	assert.GreaterOrEqual(t, len(actual), 30, fmt.Sprintf("There should be minimum like 30 stations fetched. %v stations was fetched.", len(actual)))
	assert.GreaterOrEqual(t, len(actual[5]), 70, fmt.Sprintf("Station with slice index 5 should have more sensors. Now it has %v", actual[5]))

	assert.Condition(t, func() (success bool) { return sort.StringsAreSorted(actual) }, fmt.Sprint("Should be sorted."))
}

//go:generate moq -out mocked_tests.go . IStationsCapabiltiesFetcher
type IStationsCapabiltiesFetcher interface {
	GetAllStationsCapabilities() (map[string]*AirStation, error)
}
