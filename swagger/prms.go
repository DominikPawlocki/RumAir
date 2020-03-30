package swagger

// This is used for operations that want the ID of an station in the path

// swagger:parameters getSensorsOfStation
type stationID struct {
	// The ID of the station
	//
	// in: path
	// required: true
	ID string `json:"id"`
}
