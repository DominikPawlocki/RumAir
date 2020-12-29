package swagger

import (
	airStations "github.com/dompaw/RumAir_Pmpro_Sensors_API/airStations"
	geolocalize "github.com/dompaw/RumAir_Pmpro_Sensors_API/geoLocalize"
)

// Bad request 400 response-something nok with the query string, probably
// swagger:response badRequest
type swaggBadReqResp struct {
	// in: body
	Body string
}

// Not Found 404 response
// swagger:response notFound
type swaggNotFoundResp struct {
	// in: body
	Body string
}

// Uuuuupss response
// swagger:response internalServerError
type swaggInternalSrvErrResp struct {
	// in: body
	Body string
}

// HealthCheck response
// swagger:response healthCheckResponse
type swaggSimpleResp struct {
	// in: body
	Body string
}

// Success response is a list of station sensors with all the properties.
// swagger:response sensorsResponse
type swaggGetStationSensorCodesHandlerSuccessResp struct {
	// in: body
	Body []airStations.Sensor
}

// Success response is a geolocalized station (has CitiesNearby and Lat/Lon), without its sensors (only with sensors count).
// swagger:response geolocatingStationsResponse
type swaggGeolocateStationsSuccessResp struct {
	// in: body
	Body map[string]*geolocalize.LocalizedAirStationSimplified
}

// Success response ia a list of nearest cities/districts, with a station Ids which are placed there.
// swagger:response geolocatingCitiesWithStationsResponse
type swaggGeolocateStationsByCitiesHandlerSuccessResp struct {
	// in: body
	Body map[string]*geolocalize.CitiesWithStations
}

// Success response is a list of stations indexed by ID, with its sensors with all the properties. It is pretty big JSON, MBytes probably.
// swagger:response stationsResponse
type swaggGetAllStationsCapabilitiesHandlerSuccessResp struct {
	// in: body
	Body map[string]*airStations.AirStation
}

// Success response is a geolocalized station (has CitiesNearby and Lat/Lon), with all its sensors.
// swagger:response geolocatingParticularStationResponse
type swaggGeolocateParticularStationUsingLocationIQHandlerSuccessResp struct {
	// in: body
	Body *geolocalize.LocalizedAirStation
}

// Success response is stringified list of stations and its sensor codes. Like eg : "Station : 12 can :  12HUMID_F 12HUMID_O 12PRESS_F 12PM10A"
// swagger:response allStationsSensorCodesStringifiedResponse
type swaggShowAllStationsSensorCodesHandlerSuccessResp struct {
	// in: body
	// Example:
	// [Station : 12 can : 12HUMID_F 12HUMID_O 12PRESS_F 12PM10A],
	Body []string
}

// Success response is list of sensors codes which belongs to given station
// swagger:response stationSensorCodesHandlerSuccessResp
type swaggShowStationSensorCodesHandlerSuccessResp struct {
	// in: body
	Body []string
}

// Data ordered (keyed) by sensorCode.
// swagger:response sensorDataKeyedViaCodeHandlerResponse
type swaggSensorDataKeyedViaCodeHandlerSuccessResp struct {
	// in: body
	// Example:
	//  {
	//   "data": [
	//    {
	// 		"sensorCode": "04HUMID_O:A1h",
	// 		"data": [
	// 		  {
	// 			"time": 1555632000,
	// 			"value": 55.7294
	// 		  },
	// 		  {
	// 			"time": 1555714800,
	// 			"value": 52.4889
	// 		  }
	// 		]
	// 	  },
	// 	  {
	// 		"sensorCode": "04PRESS_O:A1h",
	// 		"data": [
	// 		  {
	// 			"time": 1555632000,
	// 			"value": 1033.01
	// 		  },
	// 		  {
	// 			"time": 1555635600,
	// 			"value": 1033
	// 		  }
	// 		]
	// 	  }
	// 	]
	//   }
	Body *airStations.SensorDataKeyedViaCode
}
