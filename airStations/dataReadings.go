package airStations

import (
	"encoding/json"
	"fmt"
	"time"

	"io/ioutil"
	"net/http"
)

// https://pmpro.dacsystem.pl/webapp/data/averages?
// _dc=1597867893400&
// type=chart&
// avg=1h&
// start=1597608573&end=1597867773&
// vars=06HUMID_O%3AA1h%2C06PRESS_O%3AA1h%2C06PM10A_6_k%3AA1h%2C06PM25A_6_k%3AA1h
func FF(startTime time.Time, endTime time.Time, string sensorCode) (sensorsFetchingDataUri string) {
	var dataBaseUri = "https://pmpro.dacsystem.pl/webapp/data/averages?"
	sensorsFetchingDataUri = dataBaseUri + fmt.Sprintf("_dc=%v&type=chart&start=%v&end=%v&vars=%s", time.Now().Unix(), startTime.Unix(), endTime.Unix(), station.LatitudeSensor)
	return
}

type SingleSensorPointInTimeDataRead struct {

	// 	full pmPro answer looks like :
	//{
	// "t": 1597611600,
	// "v": 62.780000000000000,
	// "r": 62.780000000000000,
	// "s": "A",
	// "a": "d_1.9496_g_65.85_l_58.72_a_62.77997222222221_v_360_c_360_f_360_p_100",
	// "ov": 62.780000000000000,
	// "or": 0.0,
	// "os": "A",
	// "oa": "d_1.9496_g_65.85_l_58.72_a_62.77997222222221_v_360_c_360_f_360_p_100"
	//}

	Time  int64   `json:"t"`
	value float32 `json:"v"`
}

type SensorsDataInTimePeriodResponse struct {
	End    int64                               `json:"end"`
	Start  int64                               `json:"start"`
	Values [][]SingleSensorPointInTimeDataRead `json:"values"`
	Vars   []string                            `json:"vars"`
}

func GetSingleSensorDataBetweenTimePoints(fetchData IURIGetAbstracter, startTime time.Time, endTime time.Time, string sensorCode) (result SensorsDataInTimePeriodResponse, err error) {
	result = SensorsDataInTimePeriodResponse{}

	//	smt.StartDate = time.Unix(aux.StartedAt, 0)

	bytesRead, err := fetchData.DoHttpGetCall()
	if err != nil {
		return
	}

	err = json.Unmarshal(bytesRead, &result)
	if err != nil {
		fmt.Println("Error during deserializing occured. Data from `../table=Measurement&v=2`. Error is :", err)
		return
	}

	/*re := regexp.MustCompile("[0-9]+")
	result = map[string]*AirStation{} //exact same like result = make(map[string]AirStation)

	for _, measurmentType := range allMeasurments.Data {
		// Assumption ! - The 1st digits set in this string means stationId ! Like in `001NO2` the stationId is 001. Can be 2 or 3 numbers.
		stationID := re.FindAllString(measurmentType.Code, 1)[0]
		station, isExisting := result[stationID]
		if !isExisting {
			station = createNewStation(stationID)
			result[stationID] = station
		}
		if isLongitude(measurmentType.Code) {
			station.LongitudeSensor = measurmentType.Code
		}
		if isLatitude(measurmentType.Code) {
			station.LatitudeSensor = measurmentType.Code
		}
		station.Sensors = append(station.Sensors, measurmentType)
		station.SensorsCount = len(station.Sensors)
	}*/
	return
}

// kinda utils and refasctor to use this by old one !
type IURIGetAbstracter interface {
	DoHttpGetCall(uri string) ([]byte, error)
}
type URIGetAbstracter struct {
}

func (URIGetAbstracter) DoHttpGetCall(uri string) ([]byte, error) {
	return doHttpGetCall(uri)
}

func doHttpGetCall(uri string) (bytesRead []byte, err error) {
	var netResp *http.Response

	netResp, err = http.Get(uri)
	if err != nil {
		return
	}

	defer netResp.Body.Close()

	// allMeasurments slice contains whole system capability. Pretty big JSON (ca 1800 objects).
	//SLICE INITIALIZATIONS !
	//allMeasurments := make([]Sensor, 2)
	//var allMeasurments *[]Sensor = &[]Sensor{}

	bytesRead, err = ioutil.ReadAll(netResp.Body)
	if err != nil {
		fmt.Printf("Error during ReadAll bytesRead: %s err: %v. \n", bytesRead, err)
		return
	}

	fmt.Printf("%v bytes read from network for `../table=Measurement&v=2` endpoint. \n", len(bytesRead))
	return
}
