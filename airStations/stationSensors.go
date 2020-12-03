package airStations

/* The 'pmpro.dacsystem.pl' stations has sensors which collects different air measurments data - nothing unusual, uh?
But which data every station is able to collect? Does all the stations collects same data ? Does station with id X collects NO2?
The call : 'https://pmpro.dacsystem.pl/webapp/json/do?table=Measurement&v=2' answers this question,but it returns all Poland stations in one Json file.
*/

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

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

type Sensor struct {
	ID                 int     `json:"id"`
	Code               string  `json:"code"` //The first two letters of `Code` is the station Id where given station is installed! If same first 2 letters then it means same station.
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
	StartDate          uint64  `json:"start_date"`
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
func GetAllStationsCapabilities(fetchData IHttpAbstracter) (result map[string]*AirStation, err error) {
	allMeasurments := SensorsSimplifiedResponse{}

	bytesRead, err := DoHttpCallWithConsoleDots(fetchData.DoHttpGetCall, allStationsMeasurmentsURL)
	if err != nil {
		return
	}

	err = DeserializeWithConsoleDots(json.Unmarshal, bytesRead, &allMeasurments)
	if err != nil {
		fmt.Println("Error during deserializing occured. Data from `../table=Measurement&v=2`. Error is :", err)
		return
	}

	re, err := regexp.Compile("[0-9]+")
	result = map[string]*AirStation{} //exact same like result = make(map[string]AirStation)

	for _, measurmentType := range allMeasurments.Data {
		// Assumption ! - The 1st digits set in this string means stationId ! Like in `001NO2` the stationId is 001. Can be 2 or 3 numbers.
		allStationIdsFound := re.FindAllString(measurmentType.Code, 1)
		if allStationIdsFound == nil {
			fmt.Println("Not found :", measurmentType.Code)
			continue
		}
		stationID := allStationIdsFound[0]

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

func GetStationCapabilities(fetchData IHttpAbstracter, stationID string) (result *AirStation) {
	var allMeasurments *SensorsSimplifiedResponse

	bytesRead, err := DoHttpCallWithConsoleDots(fetchData.DoHttpGetCall, allStationsMeasurmentsURL)
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
func GetStationSensors(fetchData IHttpAbstracter, stationID string) (result []Sensor, err error) {
	//instead of reuturn nil - slice `zero` value default, return empty slice
	var allMeasurments *SensorsResponse

	bytesRead, err := DoHttpCallWithConsoleDots(fetchData.DoHttpGetCall, allStationsMeasurmentsURL)
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
