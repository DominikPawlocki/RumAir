package main

import (
	"fmt"

	"github.com/dompaw/RumAir/airStations"
	dataFetcher "github.com/dompaw/RumAir/dataFetcher"
	geolocalize "github.com/dompaw/RumAir/geolocalize"
)

func main() {
	fmt.Printf("Starting ...")
	//StartCron()

	sts := saveStationsCapabilitiesToFile()
	localizedStations, err := geolocalize.LocalizeStationsLocIQ(sts)
	if err != nil {
		fmt.Printf("Error during localizing occured ! %v", err)
	}
	fmt.Printf("%v stations has been localized ! \n", len(localizedStations))

	cities := geolocalize.GetStationNrPerCity(localizedStations)
	fmt.Printf("CITIES ARE : \n %s", cities)

	//server.Init()
}

func startCron() {
	cronSize, err := dataFetcher.StartCron()
	if err != nil {
		fmt.Printf("Cron NOT STARTED ! %v", err)
	}
	fmt.Printf("Cron with size %v started correctly !\n", cronSize)
}

func saveStationsCapabilitiesToFile() (result map[string]*airStations.AirStation) {
	if result = airStations.GetAllStationsCapabilities(); len(result) > 0 {
		airStations.SaveJsonToFile(result, "stationCapabilites.txt")
	}
	return
}
