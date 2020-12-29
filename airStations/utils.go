package airStations

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"sort"
	"time"
)

func DeserializeWithConsoleDots(fn func([]byte, interface{}) error, bytes []byte, result interface{}) error {
	ticker := time.NewTicker(100 * time.Millisecond)
	done := make(chan bool)

	go selectTicker(done, ticker)

	err := fn(bytes, result)

	ticker.Stop()
	done <- true
	return err
}

func selectTicker(done chan bool, ticker *time.Ticker) {
	for {
		select {
		case <-done:
			return
		case _ = <-ticker.C:
			fmt.Printf(".")
		}
	}
}

func sortAirStationsIds(stations map[string]*AirStation) (sortedKeys []string) {
	sortedKeys = make([]string, len(stations))

	var idx int = 0
	for key := range stations {
		sortedKeys[idx] = key
		idx++
	}
	sort.Strings(sortedKeys)
	return
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
