package geolocalize

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LocalizeStationsLocIQ(t *testing.T) {
	actual, _ := LocalizeStationsLocIQ(expected)

	assert.Len(t, actual, 1, fmt.Sprintf("Station with id %v should be only localized, cause it has lat and lon 'sensors'.", localizableStation.ID))
	assert.Len(t, actual["02"].CitiesNearby, 1, "It should contain 1 element like 'Rumia, Zagorze' ")
	assert.True(t, strings.HasPrefix(actual["02"].CitiesNearby[0], "Rumia"), "Should be like : Rumia, Zagórze'")
}
