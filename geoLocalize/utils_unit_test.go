package geolocalize

import (
	//"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var localizedAirStationsMock = map[string]*LocalizedAirStation{
	"02":  &LocalizedAirStation{Station: localizableStation, CitiesNearby: []string{"aaa", "bbb"}},
	"021": &LocalizedAirStation{Station: notLocalizableStation, CitiesNearby: []string{"aaa", "ccc"}}}

func Test_GetStationNrPerCity(t *testing.T) {
	result := GetStationNrPerCity(localizedAirStationsMock)

	assert.Len(t, result, 3)
	assert.Equal(t, "aaa with 2 stations : 2, 21", result[0])
	assert.Equal(t, "bbb with 1 stations : 2", result[1])
	assert.Equal(t, "ccc with 1 stations : 21", result[2])
}
