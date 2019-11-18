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
	d
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
	if measurmentTypes :=
		sensors.GetStationMeasurmentsCapabilities(sensors.SensorsToFetch["1"].ID); len(measurmentTypes) > 0 {
		sensors.SaveJsonToFile(measurmentTypes, "stationCapabilites.txt")
	}
}

func AA(){
	if sensors.GetStationIds()
}
