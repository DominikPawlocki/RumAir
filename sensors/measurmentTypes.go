package sensors

/* The stations has sensors which collects data - nothing unusual, uh?
But which data every station collects? Does all the stations collects same data ? Does station X collects NO2?
This is what this class is for.
The call : https://pmpro.dacsystem.pl/webapp/json/do?table=Measurement&v=2
returns all avialble data for all the stations, so in short it tells which station has a capability to collect what - and returns it all (not "per station").

This class finds it and returns as a result, per station */

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
	"os"
)

var measurmentURL string = "https://pmpro.dacsystem.pl/webapp/json/do?table=Measurement&v=2"

type SensorMeasurmentType struct {
	ID                 int    `json:"id"`
	Code               string `json:"code"`
	Name               string `json:"name"`
	CompoundType       string `json:"compound_type"`
	PhysicalDeviceID   int    `json:"physical_device_id"`
	PhysicalDeviceSlot string `json:"physical_device_slot"`
	UnitID             string `json:"unit_id"`
	CoefA              int    `json:"coef_a"`
	CoefB              int    `json:"coef_b"`
	TechnicalP         int    `json:"technical_p"`
	VirtualP           int    `json:"virtual_p"`
	AnalogP            int    `json:"analog_p"`
	AnalogChan         int    `json:"analog_chan"`
	BinaryP            int    `json:"binary_p"`
	BinaryChan         int    `json:"binary_chan"`
	BinaryCounter      int    `json:"binary_counter"`
	CoverageRate       int    `json:"coverage_rate"`
	AggUnit            string `json:"agg_unit"`
	Fconv              int    `json:"fconv"`
	Decimals           int    `json:"decimals"`
	Format             string `json:"format"`
	SampleType         string `json:"sample_type"`
	AverageType        string `json:"average_type"`
	Averages           string `json:"averages"`
	HighAverages       string `json:"high_averages"`
	Expression         string `json:"expression"`
	FinishDate         string `json:"finish_date"`
	IsPublished        int    `json:"is_published"`
	Timeshift          int    `json:"timeshift"`
	ManualP            int    `json:"manual_p"`
	PassiveP           int    `json:"passive_p"`
	StartDate          int    `json:"start_date"`
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

	netResp, err = http.Get(measurmentURL)
	if err != nil {
		return nil, err
	}

	defer netResp.Body.Close()

	// allMeasurments slice contains whole system capability. Pretty big JSON (ca 1800 objects).
	var allMeasurments *[]SensorMeasurmentType

	if bytesRead, err := ioutil.ReadAll(netResp.Body); err != nil && len(bytesRead) > 0 {
		err = json.Unmarshal(bytesRead, allMeasurments)

		for _, measurmentType := range *allMeasurments {
			if measurmentType.Code == stationID {
				result = append(result, measurmentType)
			}
		}
	}

	return
}

func SaveToFile(v interface{}, fileName string) (err error) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	defer f.Close()

	switch v.(type) {
	case SensorMeasurmentSimpleType:
		fmt.Fprintln(f, v.([]SensorMeasurmentSimpleType))
	case SensorMeasurmentType:
		fmt.Fprintln(f, v.([]SensorMeasurmentType))
	default:
		return errors.New("Cron not started")
	}
	f.Close()
	return
}
