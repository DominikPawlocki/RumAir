package swagger

import (
	"github.com/dompaw/RumAir_Pmpro_Sensors_API/airStations"
	geolocalize "github.com/dompaw/RumAir_Pmpro_Sensors_API/geoLocalize"
)

// No Content 204 response
// swagger:response noContent
type swaggNoContentResp struct {
	// in: body
	Body string
}

// Not Found 404 response
// swagger:response notFound
type swaggNotFoundResp struct {
	// in: body
	Body string
}

// Success response is a list of station sensors with all the properties
// swagger:response sensorsResponse
type swaggGetStationSensorCodesHandlerSuccessResp struct {
	// in: body
	Body []airStations.Sensor
}

// Uuuuupss response
// swagger:response internalServerError
type swaggInternalSrvErrResp struct {
	// in: body
	Body string
}

// Success response
// swagger:response geolocatingByGeobytesResponse
type swaggGeolocateUsingGeoBytesHandlerSuccessResp struct {
	// in: body
	//Body geolocalize.CitiesWithStations
	Body map[string]*geolocalize.LocalizedAirStationSimplified
}

// Success response
// swagger:response geolocatingCitiesWithStationsResponse
type swaggGeolocateStationsByCitiesHandlerSuccessResp struct {
	// in: body
	Body map[string]*geolocalize.CitiesWithStations
}
