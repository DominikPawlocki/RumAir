package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dompaw/RumAir/airStations"
	"github.com/gorilla/mux"
)

type MockableHTTPHandler struct {
	mockableDataFetcher airStations.IStationsCapabiltiesFetcher
	methodToBeCalled    func(w http.ResponseWriter, r *http.Request, f airStations.IStationsCapabiltiesFetcher)
}

func (m MockableHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.methodToBeCalled(w, r, m.mockableDataFetcher)
	//GetAllStationsCapabilitiesHandler(w, r, m.mockableDataFetcher)
}

func Main() {
	fmt.Println("Starting a server ..")
	myRouter := mux.NewRouter().StrictSlash(true)

	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/stations/locate/locationIQ", LocalizeAllStationsUsingLocationIQHandler)
	myRouter.HandleFunc("/stations/{id}/locate/locationIQ", LocalizeStationUsingLocationIQHandler)
	// myRouter.HandleFunc("/stations/locate/geobytes", geolocalize.LocalizeStationsGeoBytes(sts))
	// myRouter.HandleFunc("/stations/{id}/sensors", airStations.GetStationSensors(id)).Methods("POST")
	myRouter.Handle("/stations/sensors", MockableHTTPHandler{
		mockableDataFetcher: airStations.StationsCapabiltiesFetcher{},
		methodToBeCalled:    GetAllStationsCapabilitiesHandler}).Methods("GET")
	myRouter.Handle("/stations/sensors/codes", MockableHTTPHandler{
		mockableDataFetcher: airStations.StationsCapabiltiesFetcher{},
		methodToBeCalled:    ShowAllStationsSensorsCodesHandler}).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
