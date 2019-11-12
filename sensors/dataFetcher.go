package sensors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Station - reflects a physical air analyzer station put on the street or roof. It has many sensors.
//Stations differs itself, some has more sensors, some less, that its measurment capabilities differs.
type Station struct {
	ID          string
	Desc        string
	CronHandler func()
}

/*SO2
pył PM10
CO
pył PM2,5
O3
NO2
benzen

AVERAGES : averages":"A10m,A30m,A1h","high_averages":"A24h,A8h,A8h_max,A1M,A1Y"
*/

type SensorRawReadingResult struct {
	Name        string `json:"name"`
	PublicRepos int    `json:"public_repos"`
}

var SensorsToFetch = map[string]Station{
	//check it ! RU04
	"1": Station{ID: "04", Desc: "Jana III Sobieskiego", CronHandler: func() { fetchSensorDataAndSaveToDB("1573048257175") }},
	"2": Station{ID: "05", Desc: "Sabata", CronHandler: FetchSensor12345},
	"3": Station{ID: "06", Desc: "Różana", CronHandler: func() { fetchSensorDataAndSaveToDB("1573050067273") }},
	"4": Station{ID: "07", Desc: "Kujawska", CronHandler: func() { fetchSensorDataAndSaveToDB("1573050097014") }},
	"5": Station{ID: "08", Desc: "Kościelna (Skwer Plac Kaszubski)", CronHandler: func() { fetchSensorDataAndSaveToDB("1573050124901") }},
}

func fetchSensorDataAndSaveToDB(sensorID string) {
	fmt.Printf("Fetch data for sensor %s on %v \n", sensorID, time.Now())

}

func FetchSensor12345() {
	fetchSensorDataAndSaveToDB("12345")
}

func createUri(sensorID string) (url string) {
	url = fmt.Sprintf("https://api.github.com/users/%s", sensorID)
	return
	//https://pmpro.dacsystem.pl/webapp/data/averages?_dc=1571382648880&type=chart&avg=1h&start=1571123328&end=1571382528&vars=08HUMID_O%3AA1h%2C08PRESS_O%3AA1h%2C08PM10A_6_k%3AA1h%2C08PM25A_6_k%3AA1h
	//https://pmpro.dacsystem.pl/webapp/data/averages?_dc=1573496713351&type=chart&avg=1h&start=1573237393&end=1573496593&vars=05HUMID_O%3AA1h%2C05PRESS_O%3AA1h%2C05PM10A_6_k%3AA1h%2C05PM25A_6_k%3AA1h
}

func getSensorData(uri string) (res *SensorRawReadingResult, err error) {
	var netResp *http.Response
	netResp, err = http.Get(uri)
	if err != nil {
		return nil, err
	}

	defer netResp.Body.Close()
	var respBytes []byte
	res = &SensorRawReadingResult{}
	_, err = netResp.Body.Read(respBytes)
	err = json.Unmarshal(respBytes, res)

	//same like return res, err
	return
}
