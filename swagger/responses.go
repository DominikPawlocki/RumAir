package swagger

import (
	"github.com/dompaw/RumAir_Pmpro_Sensors_API/airStations"
)

// Success response
// swagger:response ok
type swaggScsResp struct {
	// in:body
	Body struct {
		// HTTP status code 200 - Status OK
		Code int `json:"code"`
		// Repository model
		Data []airStations.SensorMeasurmentType `json:"data"`
	}
}
