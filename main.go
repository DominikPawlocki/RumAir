package main

import (
	"fmt"

	"github.com/dompaw/RumAir/sensors"
)

func main() {
	stations := createAllStations()
	//sensors.AddStationsToCron(stations)
	//startCron()

	//saveStationsCapabilitiesToFile(stations)
	geoLocalizeStations(stations)
	//server.Init()
}

// AllSensors returns a slice produced from map of all sensors
func createAllStations() []sensors.Station {
	sensorsSlice := make([]sensors.Station, len(sensors.SensorsToFetch))
	idx := 0
	for _, sensor := range sensors.SensorsToFetch {
		sensorsSlice[idx] = sensor
		idx++
	}
	return sensorsSlice
}

func startCron() {
	cronSize, err := sensors.StartCron()
	if err != nil {
		fmt.Printf("Cron NOT STARTED ! %v", err)
	}
	fmt.Printf("Cron with size %v started correctly !\n", cronSize)
}

func saveStationsCapabilitiesToFile([]sensors.Station) {
	if measurmentTypes, err :=
		sensors.GetStationMeasurmentsCapabilities(sensors.SensorsToFetch["1"].ID); err == nil && len(measurmentTypes) > 0 {
		sensors.SaveJsonToFile(measurmentTypes, "stationCapabilites.txt")
	}
}

func geoLocalizeStations([]sensors.Station) {
	if _, err :=
		sensors.GetCitiesNearby(54.5708, 18.3878); err == nil {

	}
}
