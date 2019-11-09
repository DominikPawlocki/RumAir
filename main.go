package main

import (
	"fmt"

	"github.com/dompaw/RumAir/sensors"
	"github.com/dompaw/RumAir/server"
)

func main() {
	//cron := cron.New()
	sensorsSlc := AllSensors()
	for i, sensor := range sensorsSlc {
		if sensor.Handler != nil {
			sensorID, err := sensors.AddSensorToCron(i, sensor.Handler)

			if err == nil {
				fmt.Printf("Sensor %v added to Cron.\n", sensorID)
			}
		}
	}

	cronSize, err := sensors.StartCron()
	if err != nil {
		fmt.Printf("Cron NOT STARTED ! %v", err)
	}
	fmt.Printf("Cron with size %v started correctly !\n", cronSize)

	server.Init()
}

// AllSensors returns a slice produced from map of all sensors
func AllSensors() []sensors.Sensor {
	values := make([]sensors.Sensor, len(sensors.SensorsToFetch))
	idx := 0
	for _, sensor := range sensors.SensorsToFetch {
		values[idx] = sensor
		idx++
	}
	return values
}
