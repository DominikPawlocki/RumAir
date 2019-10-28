package main

import (
	"github.com/dompaw/RumAir/server"

	"github.com/robfig/cron/v3"
)

// Sensor aaaaaaaaaaaaaaa
type Sensor struct {
	ID   string
	Desc string
}

var sensors = map[string]Sensor{
	"0345391802": Sensor{ID: "0345391802", Desc: "The Hitchhiker's Guide to the Galaxy"},
	"121212":     Sensor{ID: "121212", Desc: "Cloud Native Go"},
}

func main() {
	server.Init()

	cron := cron.New()
	sensorsSlc := AllSensors()
	for i, sensor := range sensorsSlc {
		sensors.AddSensorToCron(*cron, sensor.ID, i)
	}

	cron.Start()
	defer cron.Stop()
}

// AllSensors returns a slice of all sensors
func AllSensors() []Sensor {
	values := make([]Sensor, len(sensors))
	idx := 0
	for _, sensor := range sensors {
		values[idx] = sensor
		idx++
	}
	return values
}
