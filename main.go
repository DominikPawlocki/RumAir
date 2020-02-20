package main

import (
	"fmt"

	airStations "github.com/dompaw/RumAir/airStations"
	api "github.com/dompaw/RumAir/api"
)

func main() {
	fmt.Println("Starting ...")
	api.Main()
}

func saveAllStationsCapabilitiesToFile() (result map[string]*airStations.AirStation) {
	if result, err := airStations.GetAllStationsCapabilities(airStations.StationsCapabiltiesFetcher{}); err != nil && len(result) > 0 {
		airStations.SaveJsonToFile(result, "allStationsCapabilites.txt")
	}
	return
}
