package geolocalize

import (
	"fmt"
	"github.com/dompaw/RumAir/airStations"
	"strings"
)

var geobytesBaseApiURL string = "http://getnearbycities.geobytes.com/GetNearbyCities"

type LocalizedAirStation struct {
	Station      *airStations.AirStation
	Lat          float64
	Lon          float64
	CitiesNearby []string
}

func LocalizeStationsGeoBytes(stations map[string]*airStations.AirStation) (result map[string]*LocalizedAirStation, err error) {
	result = map[string]*LocalizedAirStation{}

	for id, station := range stations {
		if station.LatitudeSensor != "" && station.LongitudeSensor != "" {
			if localizedStation, err := localizeStationGeoBytes(station); err == nil {
				result[id] = localizedStation
			}
		}
	}
	return
}

func localizeStationGeoBytes(station *airStations.AirStation) (result *LocalizedAirStation, err error) {
	result = &LocalizedAirStation{Station: station}
	if result.Lat, result.Lon, err = getStationCoordinates(station); err == nil && result.Lat != 0 && result.Lon != 0 {
		result.CitiesNearby, err = getCitiesNearbyGeoBytes(result.Lat, result.Lon)
		return
	}
	return result, fmt.Errorf("Can't localize station %v lat and lon. \n", station.ID)
}

func getReverseGeocodedCitiesGeoBytes(radius int, lat float64, lon float64) (bytesRead []byte, err error) {
	// concat strings by + not efficient but doesnt matter here
	citiesNearbyURL := geobytesBaseApiURL + fmt.Sprintf("?callback=RumAir&radius=%v&latitude=%f&longitude=%f", radius, lat, lon)
	fmt.Printf("ENDPOINT : %s. \n", citiesNearbyURL)

	return doAPIGet(citiesNearbyURL)
}

//to smaller method ! oraz inny package !
func getCitiesNearbyGeoBytes(lat float64, lon float64) (citiesNearby []string, err error) {
	radius := 0
	var reverseGeocodingStringedResponce string
	var bytesRead []byte

	for until := true; until; until = (len(reverseGeocodingStringedResponce)) < 5 {
		radius += 30 //Lets try bigger range. We need this city info, maybe there is some city further ...

		bytesRead, err = getReverseGeocodedCitiesGeoBytes(radius, lat, lon)

		if err != nil {
			fmt.Printf("Error during ReadAll bytesRead: %s err: %v. \n", bytesRead, err)
		}

		if len(bytesRead) > 0 {
			fmt.Printf("%v bytes read from network for `../getnearbycities...` endpoint for %f %f. Now, deserializing. \n", len(bytesRead), lat, lon)
			//responce is JSON-P so simple Unmarshal doesnt work here
			strs := strings.Split(string(bytesRead), ",")
			strs = removeDoubleQuotesFromSlice(strs)

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

func removeDoubleQuotesFromSlice(s []string) (result []string) {
	result = make([]string, len(s))

	for i, str := range s {
		result[i] = removeDoubleQuotes(str)
	}
	return result
}

func removeDoubleQuotes(chars string) string {
	return strings.Trim(chars, "\"")
}
