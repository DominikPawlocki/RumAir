package swagger

// swagger:parameters stationOps
type stationId struct {
	// The ID of the station
	//
	// in: path
	// required: true
	ID string
}

// swagger:parameters dateTimeOps
type day struct {
	// The day part of the date
	//
	// in: query
	// required: true
	dayNr string
}

// swagger:parameters dateTimeOps
type month struct {
	// The day part of the date
	//
	// in: query
	// required: true
	monthNr string
}
