package airStations

/* The 'pmpro.dacsystem.pl' stations has sensors which collects different air measurments data - nothing unusual, uh?
But which data every station is able to collect? Does all the stations collects same data ? Does station with id X collects NO2?
The call : 'https://pmpro.dacsystem.pl/webapp/json/do?table=Measurement&v=2' answers this question,but it returns all Poland stations in one Json file.

This code filters this huge response nicely, and it outputs to the file: which stations has which capability reported */

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var allStationsMeasurmentsURL string = "https://pmpro.dacsystem.pl/webapp/json/do?table=Measurement&v=2"

type AvailableMeasurmentsResponce struct {
	Success    bool                   `json:"success"`
	TotalCount int                    `json:"totalCount"`
	Message    string                 `json:"message"`
	Data       []SensorMeasurmentType `json:"data"`
}

type AvailableMeasurmentsSimpleResponce struct {
	TotalCount int                          `json:"totalCount"`
	Data       []SensorMeasurmentSimpleType `json:"data"`
}

// SensorMeasurmentType - describes the type of stations sensor is capable to measure with its Units, name and so on.
// The first two letters of `Code` is the station Id ! There is no specific call to fetch station Ids, so I have to deduct it from this call.
// Lets have one structure for both Unmarshall API responce and Marshall when saving to file.
// The station id field is : Code. If same first 2 letters then it means same station.
type SensorMeasurmentType struct {
	ID                 int     `json:"id"`
	Code               string  `json:"code"`
	Name               string  `json:"name"`
	CompoundType       string  `json:"compound_type"`
	PhysicalDeviceID   int     `json:"physical_device_id"`
	PhysicalDeviceSlot string  `json:"physical_device_slot"`
	UnitID             string  `json:"unit_id"`
	CoefA              float32 `json:"coef_a"`
	CoefB              float32 `json:"coef_b"`
	TechnicalP         int     `json:"technical_p"`
	VirtualP           int     `json:"virtual_p"`
	AnalogP            int     `json:"analog_p"`
	AnalogChan         int     `json:"analog_chan"`
	BinaryP            int     `json:"binary_p"`
	BinaryChan         int     `json:"binary_chan"`
	BinaryCounter      int     `json:"binary_counter"`
	CoverageRate       int     `json:"coverage_rate"`
	AggUnit            string  `json:"agg_unit"`
	Fconv              float32 `json:"fconv"`
	Decimals           int     `json:"decimals"`
	Format             string  `json:"format"`
	SampleType         string  `json:"sample_type"`
	AverageType        string  `json:"average_type"`
	Averages           string  `json:"averages"`
	HighAverages       string  `json:"high_averages"`
	Expression         string  `json:"expression"`
	FinishDate         string  `json:"finish_date"`
	IsPublished        int     `json:"is_published"`
	Timeshift          int     `json:"timeshift"`
	ManualP            int     `json:"manual_p"`
	PassiveP           int     `json:"passive_p"`
	StartDate          int     `json:"start_date"`
}

// Same like above, but simpler one
type SensorMeasurmentSimpleType struct {
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
	ID              int
	LatitudeSensor  string
	LongitudeSensor string
	Sensors         []SensorMeasurmentSimpleType
	SensorsCount    int
}

//GetAllStationsCapabilities - Stations are placed all over a Poland within `pmpro.dacsystem.pl/` system.
//This method returns all station's Ids, information if station is geolocalizable and its sensors (capabilities)
func GetAllStationsCapabilities() (result map[string]*AirStation) {
	allMeasurments := AvailableMeasurmentsSimpleResponce{}

	bytesRead := DoHttpCallWithConsoleDots(doAllMeasurmentsAPIcall)

	err := DeserializeWithConsoleDots(json.Unmarshal, bytesRead, &allMeasurments)
	if err != nil {
		fmt.Printf("Error during deserializing occured. Data from `../table=Measurement&v=2`. Error is : %v", err)
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

//GetStationSensors - Returns station all sensors.
//Returns richer sensor objects (SensorMeasurmentType) instead simpler one returned by GetAllStationsCapabilities() ...
func GetStationSensors(stationID string) (result []SensorMeasurmentType) {
	//instead of reuturn nil - slice `zero` value default, return empty slice
	var allMeasurments *AvailableMeasurmentsResponce

	bytesRead := DoHttpCallWithConsoleDots(doAllMeasurmentsAPIcall)
	err := json.Unmarshal(bytesRead, &allMeasurments)
	if err != nil {
		fmt.Printf("Error during deserializing occured. Data from `../table=Measurement&v=2`. Error is : %v", err)
		return
	}

	for _, measurmentType := range allMeasurments.Data {
		if strings.HasPrefix(measurmentType.Code, stationID) {
			result = append(result, measurmentType)
		}
	}
	fmt.Printf("Nr of results: %v", len(result))

	return allMeasurments.Data
}

func SaveJsonToFile(v interface{}, fileName string) (err error) {
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	var bytesToFile []byte

	// pattern called:  Type Assertion
	switch v.(type) {
	case []SensorMeasurmentSimpleType:
		bytesToFile, _ = json.MarshalIndent(v.([]SensorMeasurmentSimpleType), "", "\t")
	case []SensorMeasurmentType:
		bytesToFile, _ = json.MarshalIndent(v.([]SensorMeasurmentType), "", "\t")
	case map[string]*AirStation:
		bytesToFile, _ = json.MarshalIndent(v.(map[string]*AirStation), "", "\t")
	default:
		return errors.New("Saving to file : type not recognized \n")
	}
	err = ioutil.WriteFile(fileName, bytesToFile, 0644)
	return
}

func doAllMeasurmentsAPIcall() (bytesRead []byte) {
	var netResp *http.Response

	netResp, err := http.Get(allStationsMeasurmentsURL)
	if err != nil {
		return
	}

	defer netResp.Body.Close()

	// allMeasurments slice contains whole system capability. Pretty big JSON (ca 1800 objects).
	//SLICE INITIALIZATIONS !
	//allMeasurments := make([]SensorMeasurmentType, 2)
	//var allMeasurments *[]SensorMeasurmentType = &[]SensorMeasurmentType{}

	bytesRead, err = ioutil.ReadAll(netResp.Body)
	if err != nil {
		fmt.Printf("Error during ReadAll bytesRead: %s err: %v. \n", bytesRead, err)
	}

	if len(bytesRead) > 0 {
		fmt.Printf("%v bytes read from network for `../table=Measurement&v=2` endpoint. Now, deserializing. \n", len(bytesRead))
	}
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
