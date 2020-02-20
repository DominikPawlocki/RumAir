package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dompaw/RumAir_Pmpro_Sensors_API/airStations"
	"github.com/gorilla/mux"
)

type MockableHTTPHandler struct {
	mockableDataFetcher airStations.IStationsCapabiltiesFetcher
	methodToBeCalled    func(w http.ResponseWriter, r *http.Request, f airStations.IStationsCapabiltiesFetcher)
}

func (m MockableHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.methodToBeCalled(w, r, m.mockableDataFetcher)
}

func Main() {
	fmt.Println("Starting a server ..")
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/stations/locate/locationIQ", LocalizeAllStationsUsingLocationIQHandler)      //not mockable for unit testing
	myRouter.HandleFunc("/stations/{id}/locate/locationIQ", LocalizeStationUsingLocationIQHandler)     //not mockable for unit testing
	myRouter.HandleFunc("/stations/locate/geobytes", LocalizeAllStationsUsingGeoBytesHandler)          //not mockable for unit testing
	myRouter.HandleFunc("/stations/locate/locationIQ/numbersPerCity", GetStationNumbersPerCityHandler) //not mockable for unit testing

	myRouter.Handle("/stations/{id}/sensors", MockableHTTPHandler{
		mockableDataFetcher: airStations.StationsCapabiltiesFetcher{},
		methodToBeCalled:    ShowStationSensorCodesHandler}).Methods("GET")
	myRouter.Handle("/stations/sensors", MockableHTTPHandler{
		mockableDataFetcher: airStations.StationsCapabiltiesFetcher{},
		methodToBeCalled:    GetAllStationsCapabilitiesHandler}).Methods("GET")
	myRouter.Handle("/stations/sensors/codes", MockableHTTPHandler{
		mockableDataFetcher: airStations.StationsCapabiltiesFetcher{},
		methodToBeCalled:    ShowAllStationsSensorCodesHandler}).Methods("GET")

	log.Fatal(http.ListenAndServe(":80", myRouter))
}
