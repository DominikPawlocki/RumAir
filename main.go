package main

import (
	"fmt"

	airStations "github.com/dompaw/RumAir/airStations"
	api "github.com/dompaw/RumAir/api"
	dataFetcher "github.com/dompaw/RumAir/dataFetcher"
	db "github.com/dompaw/RumAir/db"
)

func main() {
	fmt.Printf("Starting ...")

	//------------------------ mongo DB simple check. to be deleted. ----------------------------------
	db.PlayMongo()

	api.Main()
}

func startCron() {
	cronSize, err := dataFetcher.StartCron()
	if err != nil {
		fmt.Printf("Cron NOT STARTED ! %v", err)
	}
	fmt.Printf("Cron with size %v started correctly !\n", cronSize)
}

func saveAllStationsCapabilitiesToFile() (result map[string]*airStations.AirStation) {
	if result, err := airStations.GetAllStationsCapabilities(airStations.StationsCapabiltiesFetcher{}); err != nil && len(result) > 0 {
		airStations.SaveJsonToFile(result, "allStationsCapabilites.txt")
	}
	return
}
