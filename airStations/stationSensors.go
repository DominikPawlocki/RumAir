package airStations

/* The 'pmpro.dacsystem.pl' stations has sensors which collects different air measurments data - nothing unusual, uh?
But which data every station is able to collect? Does all the stations collects same data ? Does station with id X collects NO2?
The call : 'https://pmpro.dacsystem.pl/webapp/json/do?table=Measurement&v=2' answers this question,but it returns all Poland stations in one Json file.
*/

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type IStationsCapabiltiesFetcher interface {
	DoAllMeasurmentsAPIcall() ([]byte, error)
}
type StationsCapabiltiesFetcher struct {
}

func (StationsCapabiltiesFetcher) DoAllMeasurmentsAPIcall() ([]byte, error) {
	return DoAllMeasurmentsAPIcall()
}

var allStationsMeasurmentsURL string = "http://pmpro.dacsystem.pl/webapp/json/do?table=Measurement&v=2"

type SensorsResponse struct {
	Success    bool     `json:"success"`
	TotalCount int      `json:"totalCount"`
	Message    string   `json:"message"`
	Data       []Sensor `json:"data"`
}

type SensorsSimplifiedResponse struct {
	TotalCount int                `json:"totalCount"`
	Data       []SensorSimplified `json:"data"`
}

/* The first two letters of `Code` is the station Id where given station is installed! If same first 2 letters then it means same station.*/
type Sensor struct {
	ID                 int       `json:"id"`
	Code               string    `json:"code"`
	Name               string    `json:"name"`
	CompoundType       string    `json:"compound_type"`
	PhysicalDeviceID   int       `json:"physical_device_id"`
	PhysicalDeviceSlot string    `json:"physical_device_slot"`
	UnitID             string    `json:"unit_id"`
	CoefA              float32   `json:"coef_a"`
	CoefB              float32   `json:"coef_b"`
	TechnicalP         int       `json:"technical_p"`
	VirtualP           int       `json:"virtual_p"`
	AnalogP            int       `json:"analog_p"`
	AnalogChan         int       `json:"analog_chan"`
	BinaryP            int       `json:"binary_p"`
	BinaryChan         int       `json:"binary_chan"`
	BinaryCounter      int       `json:"binary_counter"`
	CoverageRate       int       `json:"coverage_rate"`
	AggUnit            string    `json:"agg_unit"`
	Fconv              float32   `json:"fconv"`
	Decimals           int       `json:"decimals"`
	Format             string    `json:"format"`
	SampleType         string    `json:"sample_type"`
	AverageType        string    `json:"average_type"`
	Averages           string    `json:"averages"`
	HighAverages       string    `json:"high_averages"`
	Expression         string    `json:"expression"`
	FinishDate         string    `json:"finish_date"`
	IsPublished        int       `json:"is_published"`
	Timeshift          int       `json:"timeshift"`
	ManualP            int       `json:"manual_p"`
	PassiveP           int       `json:"passive_p"`
	StartDate          time.Time `json:"start_date"`
}

//UnmarshalJSON - is called when json.Unmarshal method executes on main type. It changes Unix timestamp from db to time.time.
func (smt *Sensor) UnmarshalJSON(data []byte) error {
	type Alias Sensor
	aux := struct {
		StartedAt int64 `json:"start_date"`
		*Alias
	}{
		Alias: (*Alias)(smt),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	smt.StartDate = time.Unix(aux.StartedAt, 0)
	return nil
}

// Same like above, but simpler one
type SensorSimplified struct {
	ID           int    `json:"id"`
	Code         string `json:"code"`
	Name         string `json:"name"`
	CompoundType string `json:"compound_type"`
	UnitID       string `json:"unit_id"`
	Decimals     int    `json:"decimals"`
	Format       string `json:"format"`
	AverageType  string `json:"average_type"`
	Averages     string `json:"averages"`
	HighAverages string `json:"high_averages"`
}

type AirStation struct {
	ID              int //ID as int doesnt make sense here, cause of eg 004
	LatitudeSensor  string
	LongitudeSensor string
	Sensors         []SensorSimplified
	SensorsCount    int
}

type AirStationSimplified struct {
	ID           int //ID as int doesnt make sense here, cause of eg 004
	SensorsCount int
}

//GetAllStationsCapabilities - Stations are placed all over a Poland within `pmpro.dacsystem.pl/` system.
//This method returns all station's Ids, information if station is geolocalizable and its sensors (capabilities)
func GetAllStationsCapabilities(fetchData IStationsCapabiltiesFetcher) (result map[string]*AirStation, err error) {
	allMeasurments := SensorsSimplifiedResponse{}

	bytesRead, err := DoHttpCallWithConsoleDots(fetchData.DoAllMeasurmentsAPIcall)
	if err != nil {
		return
	}

	err = DeserializeWithConsoleDots(json.Unmarshal, bytesRead, &allMeasurments)
	if err != nil {
		fmt.Println("Error during deserializing occured. Data from `../table=Measurement&v=2`. Error is :", err)
		return
	}

	re := regexp.MustCompile("[0-9]+")
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
	}
	return
}

func GetStationCapabilities(fetchData IStationsCapabiltiesFetcher, stationID string) (result *AirStation) {
	var allMeasurments *SensorsSimplifiedResponse

	bytesRead, err := DoHttpCallWithConsoleDots(fetchData.DoAllMeasurmentsAPIcall)
	if err != nil {
		fmt.Println("Error during getting data from `../table=Measurement&v=2`. Error is :", err)
		return
	}
	err = json.Unmarshal(bytesRead, &allMeasurments)
	if err != nil {
		fmt.Println("Error during deserializing:", err)
		return
	}
	result = createNewStation(stationID)

	for _, measurmentType := range allMeasurments.Data {
		if strings.HasPrefix(measurmentType.Code, stationID) {
			result.Sensors = append(result.Sensors, measurmentType)
			if isLongitude(measurmentType.Code) {
				result.LongitudeSensor = measurmentType.Code
			}
			if isLatitude(measurmentType.Code) {
				result.LatitudeSensor = measurmentType.Code
			}
		}
	}
	result.SensorsCount = len(result.Sensors)

	return result
}

//GetStationSensors - Returns station all sensors.
//Returns richer sensor objects (Sensor) instead simpler one returned by GetAllStationsCapabilities() ...
func GetStationSensors(fetchData IStationsCapabiltiesFetcher, stationID string) (result []Sensor, err error) {
	//instead of reuturn nil - slice `zero` value default, return empty slice
	var allMeasurments *SensorsResponse

	bytesRead, err := DoHttpCallWithConsoleDots(fetchData.DoAllMeasurmentsAPIcall)
	if err != nil {
		fmt.Println("Error during getting data from `../table=Measurement&v=2`. Error is :", err)
		return
	}
	err = json.Unmarshal(bytesRead, &allMeasurments)
	if err != nil {
		fmt.Println("Error during deserializing occured. Data from `../table=Measurement&v=2`. Error is :", err)
		return
	}

	for _, measurmentType := range allMeasurments.Data {
		if doesSensorBelongsToStation(measurmentType, stationID) {
			result = append(result, measurmentType)
		}
	}
	fmt.Println("Nr of results:", len(result))
	if len(result) == 0 {
		return nil, fmt.Errorf("No sensors for station %s found. Probably station %s doesnt exist. \n", stationID, stationID)
	}

	return
}

func ShowStationsSensorsCodes(stations map[string]*AirStation) (result []string) {
	var strBldr strings.Builder

	for _, stationID := range sortAirStationsIds(stations) {
		strBldr.Reset()
		for _, sensor := range stations[stationID].Sensors {
			strBldr.WriteString(" " + sensor.Code)
		}
		result = append(result, fmt.Sprintf("Station : %s can : %s", stationID, strBldr.String()))
	}
	return result
}

func SaveJsonToFile(v interface{}, fileName string) (err error) {
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	var bytesToFile []byte

	// pattern called:  Type Assertion
	switch v.(type) {
	case []SensorSimplified:
		bytesToFile, _ = json.MarshalIndent(v.([]SensorSimplified), "", "\t")
	case []Sensor:
		bytesToFile, _ = json.MarshalIndent(v.([]Sensor), "", "\t")
	case map[string]*AirStation:
		bytesToFile, _ = json.MarshalIndent(v.(map[string]*AirStation), "", "\t")
	default:
		return errors.New("Saving to file : type not recognized \n")
	}
	err = ioutil.WriteFile(fileName, bytesToFile, 0644)
	return
}

func DoAllMeasurmentsAPIcall() (bytesRead []byte, err error) {
	var netResp *http.Response

	netResp, err = http.Get(allStationsMeasurmentsURL)
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

//Maps and slices are reference types in Go and should be passed by values !
//also struct in Go has default value (zero value), instead of nil ! Nil for : pointers, functions, interfaces, slices, channels, and maps.

func isLatitude(code string) bool {
	if strings.Contains(code, "LAT") {
		return true
	}
	return false
}

func isLongitude(code string) bool {
	if strings.Contains(code, "LON") {
		return true
	}
	return false
}

func createNewStation(stationID string) (result *AirStation) {
	if idAsInt, err := strconv.ParseInt(stationID, 10, 64); err == nil {
		result = &AirStation{ID: int(idAsInt)}
	} else {
		result = &AirStation{ID: int(99999)}
	}
	return result
}

func doesSensorBelongsToStation(measurmentType Sensor, stationID string) bool {
	if strings.HasPrefix(measurmentType.Code, stationID) {
		re := regexp.MustCompile("[0-9]+")
		if re.FindAllString(measurmentType.Code, 1)[0] == stationID {
			return true
		}
	}
	return false
}
