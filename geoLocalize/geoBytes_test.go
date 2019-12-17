package geolocalize

import (
	"testing"

	"github.com/dompaw/RumAir/airStations"
)

func TestLocalizeStationGeoBytes(t *testing.T) {
	actual, err := LocalizeStationGeoBytes(&airStations.AirStation{
		ID:              2,
		LatitudeSensor:  "04LAT",
		LongitudeSensor: "04LON",
		//Sensors:         make([]airStations.SensorMeasurmentSimpleType, 2),
		Sensors:      []airStations.SensorMeasurmentSimpleType{
			ID: 1760,
			Code: "HES00_PA",
			"name": "PA HES061",
			"compound_type": "press",
			"unit_id": "hPa",
			"decimals": 0,
			"format": "",
			"average_type": "arithmetic",
			"averages": "A10m,A30m,A1h",
			"high_averages": "A24h,A1M,A1Y"
		},
		SensorsCount: 2})
}

//make([]SensorMeasurmentType, 2)
