package main

import (
	"fmt"

	"github.com/dompaw/RumAir/sensors"
)

func main() {
	createAllStations()
	//sensors.AddStationsToCron(stations)
	//startCron()

	//saveStationsCapabilitiesToFile(stations)
	saveStationsCapabilitiesToFile()
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

func saveStationCapabilitiesToFile() {
	if measurmentTypes :=
		sensors.GetStationSensors(sensors.SensorsToFetch["1"].ID); len(measurmentTypes) > 0 {
		sensors.SaveJsonToFile(measurmentTypes, "stationCapabilites.txt")
	}
}

func saveStationsCapabilitiesToFile() {
	if stationsWithSensors := sensors.GetAllStationsCapabilities(); len(stationsWithSensors) > 0 {
		sensors.SaveJsonToFile(stationsWithSensors, "stationCapabilites.txt")
	}
}
