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

var orig http.Handler

func Main() {
	fmt.Println("Starting a server ..")
	myRouter := mux.NewRouter().StrictSlash(true)
	//fs := http.FileServer(http.Dir("./swagger/")) //maybe jednak separated would be better, security
	//orig = http.StripPrefix("/swagger/", fs)

	//myRouter.PathPrefix("/swagger/").HandlerFunc()
	//myRouter.PathPrefix("/swagger/").Handler(fs).Methods("GET")

	myRouter.HandleFunc("/swagger/swagger.json", aa) //not mockable for unit testing

	myRouter.HandleFunc("/stations/locate/locationIQ", LocalizeAllStationsUsingLocationIQHandler)      //not mockable for unit testing
	myRouter.HandleFunc("/stations/{id}/locate/locationIQ", LocalizeStationUsingLocationIQHandler)     //not mockable for unit testing
	myRouter.HandleFunc("/stations/locate/geobytes", LocalizeAllStationsUsingGeoBytesHandler)          //not mockable for unit testing
	myRouter.HandleFunc("/stations/locate/locationIQ/numbersPerCity", GetStationNumbersPerCityHandler) //not mockable for unit testing

	// swagger:operation GET /stations/{id}/sensors stations getSensorsOfStation
	// ---
	// summary: Returns list of users by provided search parameters.
	// description: HTTP status will be returned depending on first search term (a - 400, e - 403, rest - 200)
	// parameters:
	// - name: id
	//   in: path
	//   description: search params
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/ok"
	myRouter.Handle("/stations/{id}/sensors", MockableHTTPHandler{
		mockableDataFetcher: airStations.StationsCapabiltiesFetcher{},
		methodToBeCalled:    ShowStationSensorCodesHandler}).Methods("GET")
	myRouter.Handle("/stations/sensors", MockableHTTPHandler{
		mockableDataFetcher: airStations.StationsCapabiltiesFetcher{},
		methodToBeCalled:    GetAllStationsCapabilitiesHandler}).Methods("GET")
	myRouter.Handle("/stations/sensors/codes", MockableHTTPHandler{
		mockableDataFetcher: airStations.StationsCapabiltiesFetcher{},
		methodToBeCalled:    ShowAllStationsSensorCodesHandler}).Methods("GET")

	log.Fatal(http.ListenAndServe(":5000", myRouter))
}

// ..
func aa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	fmt.Println("Fiel server !  ..")

	//fs := http.FileServer(http.Dir("./swagger/")) //maybe jednak separated would be better, security
	orig.ServeHTTP(w, r)
}

func addDefaultHeaders(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fn(w, r)
	}
}
