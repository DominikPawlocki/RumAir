package sensors

/* The 'pmpro.dacsystem.pl' stations has sensors which collects different air measurments data - nothing unusual, uh?
But which data every station is able to collect? Does all the stations collects same data ? Does station with id X collects NO2?
The call : 'https://pmpro.dacsystem.pl/webapp/json/do?table=Measurement&v=2' answers this question,but it returns all Poland stations in one Json file.

This code filters this huge response nicely, and it outputs to the file: which stations has which capability reported */

/*{"id":145,"code":"06RPM","name":"Tachometr wentylatora","compound_type":"poziom","physical_device_id":31,"physical_device_slot":"rpm","unit_id":"_","coef_a":1,
"coef_b":0,"technical_p":0,"virtual_p":0,"analog_p":0,"analog_chan":0,"binary_p":0,"binary_chan":0,"binary_counter":0,"coverage_rate":75,
"agg_unit":"_","fconv":1,"decimals":0,"format":"","sample_type":"normal","average_type":"arithmetic",
"averages":"A10m,A30m,A1h","high_averages":"A24h,A8h,A8h_max,A1M,A1Y","expression":"","finish_date":"",
"is_published":0,"timeshift":0,"manual_p":0,"passive_p":0,"start_date":1449754401}*/

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
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

//SensorMeasurmentType - lets have ONE structure for both Unmarshall API responce and Marshall when saving to file.
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

func GetStationMeasurmentsCapabilities(stationID string) (result []SensorMeasurmentType, err error) {
	var netResp *http.Response

	netResp, err = http.Get(allStationsMeasurmentsURL)
	if err != nil {
		fmt.Printf("Error during asking endpoint %s %v.", allStationsMeasurmentsURL, err)
		return nil, err
	}

	defer netResp.Body.Close()

	// allMeasurments slice contains whole system capability. Pretty big JSON (ca 1800 objects).
	//SLICE INITIALIZATIONS !
	//allMeasurments := make([]SensorMeasurmentType, 2)
	//var allMeasurments *[]SensorMeasurmentType = &[]SensorMeasurmentType{}
	var allMeasurments AvailableMeasurmentsResponce

	bytesRead, err := ioutil.ReadAll(netResp.Body)
	if err != nil {
		fmt.Printf("Error during ReadAll bytesRead: %s err: %v. \n", bytesRead, err)
	}

	if len(bytesRead) > 0 {
		fmt.Printf("%v bytes read from network for `../table=Measurement&v=2` endpoint for %s. Now, deserializing. \n", len(bytesRead), stationID)
		err = json.Unmarshal(bytesRead, &allMeasurments)
		if err != nil {
			fmt.Printf("Error during deserializing station %s occured. Data from `../table=Measurement&v=2`. Error is : %v", stationID, err)
			return nil, err
		}
		for _, measurmentType := range allMeasurments.Data {
			if strings.HasPrefix(measurmentType.Code, stationID) {
				result = append(result, measurmentType)
			}
		}
	}
	fmt.Printf("Nr of results: %v", len(result))
	return
}

func SaveToFile(v interface{}, fileName string) (err error) {
	//f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	//defer f.Close()

	var bytesToFile []byte

	switch v.(type) {
	case []SensorMeasurmentSimpleType:
		bytesToFile, _ = json.MarshalIndent(v.([]SensorMeasurmentSimpleType), "", "\t")
	case []SensorMeasurmentType:
		bytesToFile, _ = json.MarshalIndent(v.([]SensorMeasurmentType), "", "\t")
	default:
		return errors.New("Saving to file : type not recognized \n")
	}
	err = ioutil.WriteFile(fileName, bytesToFile, 0644)
	//f.Close()
	return
}
