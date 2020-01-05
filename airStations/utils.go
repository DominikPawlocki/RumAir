package airStations

import (
	"fmt"
	"sort"
	"time"
)

func DoHttpCallWithConsoleDots(fn func() ([]byte, error)) (bytesRead []byte, err error) {
	ticker := time.NewTicker(100 * time.Millisecond)
	done := make(chan bool)

	go selectTicker(done, ticker)

	bytesRead, err = fn()

	ticker.Stop()
	done <- true
	return
}

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
