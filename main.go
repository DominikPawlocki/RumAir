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
		if sensor.CronHandler != nil {
			sensorID, err := sensors.AddSensorToCron(i, sensor.CronHandler)

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
	sensorsSlice := make([]sensors.Sensor, len(sensors.SensorsToFetch))
	idx := 0
	for _, sensor := range sensors.SensorsToFetch {
		sensorsSlice[idx] = sensor
		idx++
	}
	return sensorsSlice
}
