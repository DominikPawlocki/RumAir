// Package classification AbaxTest..
//
// Hi !
//
// Please take a look at this app as a Hackerrank task done little bit wider.
//
// It is contenerized application, consist of two containers, public swaggerUI (pulled from DockerHub) and custom AbaxTestApi one.
// Both are deployed on Azure Docker Instances service.
//
// My container (AbaxTestApi) is build via Dockerfile and then its pushed to my private Azure Docker Container by Azure DevOps pipeline YML (./'az-pipelines.docker.yml )
// During this Docker build, there is a 'swagger.json' artifact produced also, validated and stored to be static hosted at /swagger/swagger.json endpoint.
//
// The app is in Golang, which I am into intesively last times, but it doesnt matter. I just wanted to practice it, I could do it in C# also.
//
// Still almost non of Unit tests, some parts are fixed instead to be injected during deploy, and it wasnt done in 1 hour :-) but ...
// at least I learnt a lot, and I hope you like it.
//
//	Schemes: http
//	BasePath:
//	Version: new attempt :-)
//	License: MIT http://opensource.org/licenses/MIT
//	Contact: Dominik Pawlocki <d.pawlocki@onet.pl> https://github.com/DominikPawlocki/AbaxTest1
//	Host: abaxtestapi.northeurope.azurecontainer.io:5000
//	Description: Abax test
//
// Still - good task! Thx and have a good day !
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
