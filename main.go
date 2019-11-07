package main

import (
	"fmt"

	"github.com/dompaw/RumAir/sensors"
	"github.com/dompaw/RumAir/server"
)

// Sensor aaaaaaaaaaaaaaa
type Sensor struct {
	ID   string
	Desc string
}

var sensorsToFetch = map[string]Sensor{
	"1573048257175": Sensor{ID: "1573048257175", Desc: "Jana III Sobieskiego"},
	"1573050028266": Sensor{ID: "1573050028266", Desc: "Sabata"},
	"1573050067273": Sensor{ID: "1573050067273", Desc: "Różana"},
	"1573050097014": Sensor{ID: "1573050097014", Desc: "Kujawska"},
	"1573050124901": Sensor{ID: "1573050124901", Desc: "Kościelna (Skwer Plac Kaszubski)"},
}

SO2
pył PM10
CO
pył PM2,5
O3
NO2
benzen

func main() {
	//cron := cron.New()
	sensorsSlc := AllSensors()
	for i, sensor := range sensorsSlc {
		sensorID, err := sensors.AddSensorToCron(sensor.ID, i)
		if err == nil {
			fmt.Printf("Sensor %v added to Cron.\n", sensorID)
		}
	}

	/*if (cronSize, err) := sensors.StartCron(); err != nil {
		return nil, err
	}*/

	cronSize, err := sensors.StartCron()
	if err != nil {
		fmt.Printf("Cron NOT STARTED ! %v", err)
	}
	fmt.Printf("Cron with size %v started correctly !\n", cronSize)

	server.Init()
}

// AllSensors returns a slice of all sensors
func AllSensors() []Sensor {
	values := make([]Sensor, len(sensorsToFetch))
	idx := 0
	for _, sensor := range sensorsToFetch {
		values[idx] = sensor
		idx++
	}
	return values
}
