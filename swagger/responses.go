package swagger

import (
	"github.com/dompaw/RumAir_Pmpro_Sensors_API/airStations"
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

// Success response
// swagger:response sensorsResponse
type swaggShowStationSensorCodesHandlerSuccessResp struct {
	// in: body
	Body []airStations.SensorMeasurmentType
}

// InternalServerError response
// swagger:response internalServerError
type swaggInternalSrvErrResp struct {
	// in: body
	Body string
}
