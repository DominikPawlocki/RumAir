package sensors

import (
	"fmt"
	"time"
)

// Sensor aaaaaaaaaaaaaaa
type Sensor struct {
	ID      string
	Desc    string
	Handler func()
}

var SensorsToFetch = map[string]Sensor{
	"1573048257175": Sensor{ID: "1573048257175", Desc: "Jana III Sobieskiego", Handler: func() { fetchSensorDataAndSaveToDB("1573048257175") }},
	"1573050028266": Sensor{ID: "1573050028266", Desc: "Sabata", Handler: FetchSensor12345},
	"1573050067273": Sensor{ID: "1573050067273", Desc: "Różana", Handler: func() { fetchSensorDataAndSaveToDB("1573050067273") }},
	"1573050097014": Sensor{ID: "1573050097014", Desc: "Kujawska", Handler: func() { fetchSensorDataAndSaveToDB("1573050097014") }},
	"1573050124901": Sensor{ID: "1573050124901", Desc: "Kościelna (Skwer Plac Kaszubski)", Handler: func() { fetchSensorDataAndSaveToDB("1573050124901") }},
}

func fetchSensorDataAndSaveToDB(sensorID string) {
	fmt.Printf("Fetch data for sensor %s on %v \n", sensorID, time.Now())
}

func FetchSensor12345() {
	fetchSensorDataAndSaveToDB("12345")
}

/*SO2
pył PM10
CO
pył PM2,5
O3
NO2
benzen*/
