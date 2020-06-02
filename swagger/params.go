package swagger

// swagger:parameters sensorsFetching
type id struct {
	// The ID of the station
	//
	// in: path
	// required: true
	ID string
}

// swagger:parameters sensorsFetching
type deviceParentId struct {
	// The ID of the station
	//
	// in: query
	// required: false
	ID int32
}
