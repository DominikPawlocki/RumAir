// Package classification RumAir..
//
// PMPRO system API
//
// It is contenerized application, consist of two containers, public swaggerUI (pulled from DockerHub) and custom Golang API one.
// Both are deployed on Azure Docker Instances service.
//
//
//	Schemes: http
//	Version: 1.0
//	License: MIT http://opensource.org/licenses/MIT
//	Contact: Dominik Pawlocki <d.pawlocki@onet.pl> https://github.com/DominikPawlocki
//	Description: RumAir API
//
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
