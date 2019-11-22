/* The 'pmpro.dacsystem.pl' system has air monitoring stations in several cities all over a Poland.
Cities which uses this system usually has webpages exposed like 'rumia.powietrze.eu', monitoring current air status with short (3 days) history.
But there is no list of a pages like that, and a call 'https://pmpro.dacsystem.pl/webapp/json/do?table=Measurement&v=2' reveals there are like 60 stations over the Poland.
The question is : which stations are where ?

This code answers this question, using public geocoding API from  Geobytes : `https://geobytes.com/get-nearby-cities-api/`.
Their API is mostly used for geolocating IP adresses probably, but also has a possibility to find nearest city by latitutude/longitude .
And...
'Pmpro' system stations has lat/long coordinates exposed !

This code outputs to the file, which station (id) is located where (nearest city).
Thats it ! Now I know which stations nearby my place Im interrested in grabbing history from !

*/

package sensors

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type LocalizedAirStation struct {
	Station      *AirStation
	Lat          float64
	Lon          float64
	CitiesNearby []string
}

var geoBytesBaseApiURL string = "http://getnearbycities.geobytes.com/GetNearbyCities"
var pmproSystemBaseApiURL string = "https://pmpro.dacsystem.pl/webapp/data"

func LocalizeStations(stations map[string]*AirStation) (result map[string]*LocalizedAirStation, err error) {
	result = map[string]*LocalizedAirStation{}

	for _, station := range stations {
		if localizedStation, err := LocalizeStation(station); err != nil {
			result[string(station.ID)] = localizedStation //&LocalizedAirStation{Station: station, Lat: localizedStation.}
		}
	}
	return
}

func LocalizeStation(station *AirStation) (result *LocalizedAirStation, err error) {
	if station.LatitudeSensor != "" && station.LongitudeSensor != "" {
		result = &LocalizedAirStation{Station: station}

		if result.Lat, result.Lon, err = getStationLocation(station); err == nil && result.Lat != 0 && result.Lon != 0 {
			result.CitiesNearby, err = getCitiesNearby(result.Lat, result.Lon)
			if err != nil {
				return
			}
		}
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

		bytesRead, err = getReverseGeocodedCities(radius, lat, lon)

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
				} else return citiesNearby, err
				
			}
		}
	}
	return
}

func getReverseGeocodedCities(radius int, lat float64, lon float64) (bytesRead []byte, err error) {
	// concat strings by + not efficient but doesnt matter here
	citiesNearbyURL := geoBytesBaseApiURL + fmt.Sprintf("?callback=RumAir&radius=%v&latitude=%f&longitude=%f", radius, lat, lon)

	return doAPIGet(citiesNearbyURL)
}

func getStationLocation(station *AirStation) (latitude float64, longitude float64, err error) {
	//https://pmpro.dacsystem.pl/webapp/data/averages?type=chart&start=1561939200&end=1561949200&vars=27LON
	curr, currMinus2h := nowAndMInus2hInUnixTimestamp()

	stationLatitudeURI := pmproSystemBaseApiURL + fmt.Sprintf("/averages?type=chart&start=%v&end=%v&vars=%s", currMinus2h, curr, station.LatitudeSensor)
	latitude, err = getLatOrLonFromAPI(stationLatitudeURI)

	stationLongitudeURI := pmproSystemBaseApiURL + fmt.Sprintf("/averages?type=chart&start=%v&end=%v&vars=%s", currMinus2h, curr, station.LongitudeSensor)
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

	fmt.Printf("ENDPOINT : %s. \n", sensorCallURI)

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
		/*result, err = strconv.ParseFloat(resp.Values[0][0].V, 64)
		if err != nil {
			fmt.Printf("Error during parsing string to float %s %v.\n", resp.Values[0][0].V, err)
			return
		}*/
	}
	return
}

func nowAndMInus2hInUnixTimestamp() (current int64, currentMinus2h int64) {
	current = time.Now().Unix()
	currentMinus2h = time.Now().Add(-3 * time.Hour).Unix()

	return
}
