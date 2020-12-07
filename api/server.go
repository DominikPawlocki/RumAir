package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dompaw/RumAir_Pmpro_Sensors_API/airStations"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type MockableHTTPHandler struct {
	mockableDataFetcher airStations.IHttpAbstracter
	methodToBeCalled    func(w http.ResponseWriter, r *http.Request, f airStations.IHttpAbstracter)
}

func (m MockableHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.methodToBeCalled(w, r, m.mockableDataFetcher)
}

var fileServerHandler http.Handler

func Main() {
	fmt.Println("Starting a server ..")
	myRouter := mux.NewRouter().StrictSlash(true)

	fileServerHandler = http.StripPrefix("/swagger/", http.FileServer(http.Dir("/tmp/swagger/")))
	myRouter.HandleFunc("/swagger/swagger.json", fileServerHandler.ServeHTTP) //not mockable for unit testing

	myRouter.HandleFunc("/stations/locate/locationIQ", LocalizeAllStationsUsingLocationIQHandler).Methods("GET")      //not mockable for unit testing
	myRouter.HandleFunc("/stations/{id}/locate/locationIQ", LocalizeStationUsingLocationIQHandler).Methods("GET")     //not mockable for unit testing
	myRouter.HandleFunc("/stations/locate/geobytes", LocalizeAllStationsUsingGeoBytesHandler).Methods("GET")          //not mockable for unit testing
	myRouter.HandleFunc("/stations/locate/locationIQ/numbersPerCity", GetStationNumbersPerCityHandler).Methods("GET") //not mockable for unit testing

	myRouter.Handle("/stations/{id:[0-9]+}/sensors", MockableHTTPHandler{
		mockableDataFetcher: airStations.HttpAbstracter{},
		methodToBeCalled:    GetStationSensorsHandler}).Methods("GET")
	myRouter.Handle("/stations/sensors", MockableHTTPHandler{
		mockableDataFetcher: airStations.HttpAbstracter{},
		methodToBeCalled:    GetAllStationsCapabilitiesHandler}).Methods("GET")
	myRouter.Handle("/stations/sensors/codes", MockableHTTPHandler{
		mockableDataFetcher: airStations.HttpAbstracter{},
		methodToBeCalled:    ShowAllStationsSensorCodesHandler}).Methods("GET")

	myRouter.Path("/stations/{stationId}/data").
		Queries("day", "{day}", "month", "{month}", "year", "{year}", "sensorCodes", "{sensorCodes}").
		HandlerFunc(func(h1 http.ResponseWriter, h2 *http.Request) { AAAAAAAAAAAAA(h1, h2, airStations.HttpAbstracter{}) })

	//myRouter.Queries("version", "{version}").Name("item/version").HandlerFunc("/stations/{stationId}/data",

	/*Handle("/stations/{stationId}/data", MockableHTTPHandler{
	mockableDataFetcher: airStations.HttpAbstracter{},
	methodToBeCalled:    AAAAAAAAAAAAA}).Methods("GET")*/

	myRouter.HandleFunc("/healthCheck", healthCheck).Methods("GET")

	// all origins accepted with simple methods (GET, POST). Security antipattern, will look there later.
	handlerWithAllCorsEnabled := cors.Default().Handler(myRouter)

	log.Fatal(http.ListenAndServe(":5000", handlerWithAllCorsEnabled))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /healthCheck healthCheck basicPingPongHealthCheck
	// Ping ? Pong .
	// ---
	// produces:
	// - application/json
	// responses:
	//   "200":
	//     "$ref": "#/responses/healthCheckResponse"
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to Rumia air monitoring system.")
}
