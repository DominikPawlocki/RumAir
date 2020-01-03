package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Main() {
	fmt.Println("Starting a server ..")
	myRouter := mux.NewRouter().StrictSlash(true)

	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/stations/locate/locationIQ", LocalizeAllStationsUsingLocationIQHandler)
	myRouter.HandleFunc("/stations/{id}/locate/locationIQ", LocalizeStationUsingLocationIQHandler)
	// myRouter.HandleFunc("/stations/locate/geobytes", geolocalize.LocalizeStationsGeoBytes(sts))
	// myRouter.HandleFunc("/stations/{id}/sensors", airStations.GetStationSensors(id)).Methods("POST")
	myRouter.HandleFunc("/stations/sensors", GetAllStationsCapabilitiesHandler).Methods("GET")
	myRouter.HandleFunc("/stations/sensors/simplified", ShowAllStationsSensorsCodesHandler).Methods("GET")

	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
