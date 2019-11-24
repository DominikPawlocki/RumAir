package main

import (
	"fmt"

	"github.com/dompaw/RumAir/sensors"
)

func main() {
	createAllStations()
	//sensors.AddStationsToCron(stations)
	//startCron()

	//saveStationsCapabilitiesToFile()

	sts := sensors.GetAllStationsCapabilities()
	localizedStations, err := sensors.LocalizeStations(sts)
	if err != nil {
		fmt.Printf("Error during localizing occured ! %v", err)
	}
	fmt.Printf("%v stations has been localized ! \n", len(localizedStations))

	cities := sensors.GetStationNrPerCity(localizedStations)
	fmt.Printf("CITIES ARE : %s \n", cities)
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

func saveStationsCapabilitiesToFile() {
	if stationsWithSensors := sensors.GetAllStationsCapabilities(); len(stationsWithSensors) > 0 {
		sensors.SaveJsonToFile(stationsWithSensors, "stationCapabilites.txt")
	}
}
