package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Init() {
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Fatal error, API server NOT started ! %v", err)
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Starting a server ..")
	myRouter := mux.NewRouter().StrictSlash(true)

	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/stations/locate/locationIQ", LocalizeAllStationsUsingLocationIQHandler)
	myRouter.HandleFunc("/stations/{id}/locate/locationIQ", LocalizeStationUsingLocationIQHandler)
	// myRouter.HandleFunc("/stations/locate/geobytes", geolocalize.LocalizeStationsGeoBytes(sts))
	// myRouter.HandleFunc("/stations/{id}/sensors", airStations.GetStationSensors(id)).Methods("POST")
	// myRouter.HandleFunc("/stations/sensors", airStations.GetAllStationsCapabilities()).Methods("POST")
	myRouter.HandleFunc("/stations/sensors/simplified", ShowStationsSensorsCodesHandler).Methods("GET")

	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
