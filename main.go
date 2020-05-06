// Package classification RumAir API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//	Schemes: http
//	BasePath: /v2
//	Version: 1.0.1
//	License: MIT http://opensource.org/licenses/MIT
//	Contact: Dmnk <domin@wp.eu> https://githublink.com
//	Host: rumair.azurewebsites.net
//
//	Consumes:
//		- application/json
//
//	Produces:
//		- application/json
//
// swagger:meta
package main

import (
	"fmt"

	airStations "github.com/dompaw/RumAir_Pmpro_Sensors_API/airStations"
	api "github.com/dompaw/RumAir_Pmpro_Sensors_API/api"

	_ "github.com/dompaw/RumAir_Pmpro_Sensors_API/swagger"
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
