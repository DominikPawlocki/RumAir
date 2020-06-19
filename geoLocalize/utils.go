package geolocalize

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/dompaw/RumAir_Pmpro_Sensors_API/airStations"
)

type LocalizedAirStation struct {
	Station      *airStations.AirStation
	Lat          float64
	Lon          float64
	CitiesNearby []string
}

type LocalizedAirStationSimplified struct {
	Station      *airStations.AirStationSimplified
	Lat          float64
	Lon          float64
	CitiesNearby []string
}

type CityWithStations struct {
	City             string
	Count            int
	StationIdsConcat string
}

type CitiesWithStations struct {
	Localized              []*CityWithStations
	WereLocalizedCount     int
	NotAbleToLocalizeCount int
}

func GetStationNrPerCity(localized map[string]*LocalizedAirStation) (result []*CityWithStations) {
	type stationsPerCity struct {
		StationIdsConcat string
		Count            int
	}
	var strBldr strings.Builder
	citiesNoDuplicates := map[string]*stationsPerCity{}

	for _, sts := range localized {
		for _, city := range sts.CitiesNearby {
			if spc, ok := citiesNoDuplicates[city]; !ok {
				citiesNoDuplicates[city] = &stationsPerCity{StationIdsConcat: strconv.Itoa(sts.Station.ID), Count: 1}
			} else {
				spc.Count++
				spc.StationIdsConcat += fmt.Sprintf("%s %s", ",", strconv.Itoa(sts.Station.ID))
			}
		}
	}

	var keys []string = make([]string, len(citiesNoDuplicates))
	itr := 0
	for i := range citiesNoDuplicates {
		keys[itr] = i
		itr++
	}
	sort.Strings(keys)
	result = make([]*CityWithStations, len(keys))

	for o, city := range keys {
		strBldr.WriteString(citiesNoDuplicates[city].StationIdsConcat)
		result[o] = &CityWithStations{City: city,
			Count:            citiesNoDuplicates[city].Count,
			StationIdsConcat: strBldr.String()}
		strBldr.Reset()
	}

	return
}

func getLatOrLonFromAPI(sensorCallURI string) (result float64, err error) {
	type LimitedOneValueResponse struct {
		End    int `json:"end"`
		Start  int `json:"start"`
		Values [][]struct {
			T int     `json:"t"`
			V float64 `json:"v"`
		} `json:"values"`
		Vars []string `json:"vars"`
	}

	resp := &LimitedOneValueResponse{}
	bytesRead, err := doAPIGet(sensorCallURI)
	if err != nil {
		fmt.Printf("Error during asking endpoint %s %v.\n", sensorCallURI, err)
		return
	}
	if len(bytesRead) == 0 {
		fmt.Printf("0 bytes recieved from endpoint %s.\n", sensorCallURI)
		return 0, fmt.Errorf("0 bytes recieved from endpoint %s", sensorCallURI)
	}
	err = json.Unmarshal(bytesRead, resp)
	if err != nil {
		fmt.Printf("Error during Unmarshall API responce %v.\n", err)
		return
	}

	if resp.Values[0] != nil && len(resp.Values[0]) > 0 && resp.Values[0][0].V != 0 {
		result = resp.Values[0][0].V
		fmt.Printf(" %v for : %s. \n", result, sensorCallURI)
	}
	return
}

//'Pmpro' system stations has lat/long coordinates exposed ! http://pmpro.dacsystem.pl/webapp/data/averages?type=chart&start=1561939200&end=1561949200&vars=27LON
func getStationCoordinates(station *airStations.AirStation) (latitude float64, longitude float64, err error) {
	var pmproSystemBaseAPIURL string = "http://pmpro.dacsystem.pl/webapp/data"

	curr, currMinus2h := nowAndMInus2hInUnixTimestamp()

	stationLatitudeURI := pmproSystemBaseAPIURL + fmt.Sprintf("/averages?type=chart&start=%v&end=%v&vars=%s", currMinus2h, curr, station.LatitudeSensor)
	latitude, err = getLatOrLonFromAPI(stationLatitudeURI)

	stationLongitudeURI := pmproSystemBaseAPIURL + fmt.Sprintf("/averages?type=chart&start=%v&end=%v&vars=%s", currMinus2h, curr, station.LongitudeSensor)
	longitude, err = getLatOrLonFromAPI(stationLongitudeURI)

	return
}

func doAPIGet(uri string) (bytesRead []byte, err error) {
	var netResp *http.Response

	netResp, err = http.Get(uri)
	if err != nil {
		fmt.Printf("Error during asking endpoint %s %v.", uri, err)
		return nil, err
	}

	defer netResp.Body.Close()
	bytesRead, err = ioutil.ReadAll(netResp.Body)

	return
}

func nowAndMInus2hInUnixTimestamp() (current int64, currentMinus2h int64) {
	current = time.Now().Unix()
	currentMinus2h = time.Now().Add(-3 * time.Hour).Unix()

	return
}

func SleepWithOutputDotsOnConsole(d time.Duration) {
	ticker := time.NewTicker(200 * time.Millisecond)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case _ = <-ticker.C:
				fmt.Printf(".")
			}
		}
	}()

	time.Sleep(d)

	ticker.Stop()
	done <- true
}
