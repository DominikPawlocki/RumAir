package geolocalize

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/dompaw/RumAir/airStations"
)

var geobytesBaseApiURL string = "http://getnearbycities.geobytes.com/GetNearbyCities"

type LocalizedAirStation struct {
	Station      *airStations.AirStation
	Lat          float64
	Lon          float64
	CitiesNearby []string
}

func LocalizeStations(stations map[string]*airStations.AirStation) (result map[string]*LocalizedAirStation, err error) {
	result = map[string]*LocalizedAirStation{}

	for id, station := range stations {
		if station.LatitudeSensor != "" && station.LongitudeSensor != "" {
			if localizedStation, err := LocalizeStation(station); err == nil {
				result[id] = localizedStation
			}
		}
	}
	return
}

func LocalizeStation(station *airStations.AirStation) (result *LocalizedAirStation, err error) {
	result = &LocalizedAirStation{Station: station}
	if result.Lat, result.Lon, err = getStationCoordinates(station); err == nil && result.Lat != 0 && result.Lon != 0 {
		result.CitiesNearby, err = getCitiesNearby(result.Lat, result.Lon)
		return
	}
	return result, fmt.Errorf("Can't localize station %v lat and lon. \n", station.ID)
}

func GetStationNrPerCity(localized map[string]*LocalizedAirStation) string {
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
				spc.StationIdsConcat += fmt.Sprintf("%s %s ", ",", strconv.Itoa(sts.Station.ID))
			}
		}
	}

	for city, nrOfStations := range citiesNoDuplicates {
		strBldr.WriteString(city)
		strBldr.WriteString(" with ")
		strBldr.WriteString(strconv.Itoa(nrOfStations.Count))
		strBldr.WriteString(" stations : ")
		strBldr.WriteString(nrOfStations.StationIdsConcat)
		strBldr.WriteString("\n")
	}

	return strBldr.String()
}

func getReverseGeocodedCitiesGeobytes(radius int, lat float64, lon float64) (bytesRead []byte, err error) {
	// concat strings by + not efficient but doesnt matter here
	citiesNearbyURL := geobytesBaseApiURL + fmt.Sprintf("?callback=RumAir&radius=%v&latitude=%f&longitude=%f", radius, lat, lon)
	fmt.Printf("ENDPOINT : %s. \n", citiesNearbyURL)

	return doAPIGet(citiesNearbyURL)
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
		fmt.Printf("%v for : %s. \n", result, sensorCallURI)
	}
	return
}

//to smaller method ! oraz inny package !
func getCitiesNearby(lat float64, lon float64) (citiesNearby []string, err error) {
	radius := 0
	var reverseGeocodingStringedResponce string
	var bytesRead []byte

	for until := true; until; until = (len(reverseGeocodingStringedResponce)) < 5 {
		radius += 30 //Lets try bigger range. We need this city info, maybe there is some city further ...

		bytesRead, err = getReverseGeocodedCitiesGeobytes(radius, lat, lon)

		if err != nil {
			fmt.Printf("Error during ReadAll bytesRead: %s err: %v. \n", bytesRead, err)
		}

		if len(bytesRead) > 0 {
			fmt.Printf("%v bytes read from network for `../getnearbycities...` endpoint for %f %f. Now, deserializing. \n", len(bytesRead), lat, lon)
			//responce is JSON-P so simple Unmarshal doesnt work here
			strs := strings.Split(string(bytesRead), ",")
			if len(strs) > 1 {
				//slice is immutable - append is good enough here, but might me bottleneck in different situation
				citiesNearby = append(citiesNearby, strs[1])
				if len(strs) > 15 {
					//second city also
					citiesNearby = append(citiesNearby, strs[14])
					return
				}
				return
			}
		}
	}
	return
}
