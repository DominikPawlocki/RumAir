package geolocalize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
