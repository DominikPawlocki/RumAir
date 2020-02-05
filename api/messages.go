package api

var stationsCapabilitesFetchingError = "Cannot fetch stations. \n"
var deserializingSensorsPerStationError = "Error on deserializing sensorsPerStation, when stations seems fetched. \n"
var emptySensorsPerStationError = "Empty result Sensors per stations, when stations seems fetched. \n"
var locationIQdeserializingError = "Error during serializing response from geolocalizing stations via LocationIQ API. \n"
var geoBytesdeserializingError = "Error during serializing response from geolocalizing stations via GeoBytes API. \n"
var locationIQfetchingError = "Cannnot fetch location from LocationIQ service. \n"
var geoBytesfetchingError = "Cannnot fetch location from GeoByes service. \n"

// Common errors
/*var (
	ErrNegSqrt    = errors.New("sqrt of negative number")
	ErrNoSolution = errors.New("no solution found")
)*/
