package main

import (
	"fmt"

	"github.com/dompaw/RumAir/airStations"
	geolocalize "github.com/dompaw/RumAir/geolocalize"
)

func main() {
	fmt.Printf("Starting ...")
	createAllStations()
	//sensors.AddStationsToCron(stations)
	//startCron()

	//saveStationsCapabilitiesToFile()

	sts := airStations.GetAllStationsCapabilities()
	localizedStations, err := geolocalize.LocalizeStationsLocIQ(sts)
	if err != nil {
		fmt.Printf("Error during localizing occured ! %v", err)
	}
	fmt.Printf("%v stations has been localized ! \n", len(localizedStations))

	cities := geolocalize.GetStationNrPerCity(localizedStations)
	fmt.Printf("CITIES ARE : \n %s", cities)
	//server.Init()
}

// AllSensors returns a slice produced from map of all sensors
func createAllStations() []airStations.Station {
	airStationsSlice := make([]airStations.Station, len(airStations.SensorsToFetch))
	idx := 0
	for _, sensor := range airStations.SensorsToFetch {
		airStationsSlice[idx] = sensor
		idx++
	}
	return airStationsSlice
}

func startCron() {
	cronSize, err := airStations.StartCron()
	if err != nil {
		fmt.Printf("Cron NOT STARTED ! %v", err)
	}
	fmt.Printf("Cron with size %v started correctly !\n", cronSize)
}

func saveStationsCapabilitiesToFile() {
	if stationsWithSensors := airStations.GetAllStationsCapabilities(); len(stationsWithSensors) > 0 {
		airStations.SaveJsonToFile(stationsWithSensors, "stationCapabilites.txt")
	}
}
