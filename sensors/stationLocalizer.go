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

	//"errors"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type LocalizedAirStation struct {
	Station      *AirStation
	Lat          string
	Lon          string
	CitiesNearby []string
}

var geoBytesBaseApiURL string = "http://getnearbycities.geobytes.com/GetNearbyCities"
var pmproSystemBaseApiURL string = "https://pmpro.dacsystem.pl/webapp/data"

/*func GetStationLonLat(station *AirStation) (result *LocalizedAirStation) {
	if station.LatitudeSensor && station.LongitudeSensor {


		if len(bytesRead) > 0 {
			err = json.Unmarshal(bytesRead, res *odpowiedzLATLON)

			if(err==null){
				station.
			}
		}
	}
}*/

//to smaller method ! oraz inny package !
func GetCitiesNearby(lat float32, lon float32) (citiesNearby []string, err error) {
	radius := 0
	var reverseGeocodingStringedResponce string
	var bytesRead []byte

	for until := true; until; until = (len(reverseGeocodingStringedResponce)) < 5 {
		//Lets try bigger range !
		radius += 30

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
				}
			}
		}
		return
	}
	return
}

func getReverseGeocodedCities(radius int, lat float32, lon float32) (bytesRead []byte, err error) {
	// concat strings by + not efficient but doesnt matter here
	citiesNearbyURL := geoBytesBaseApiURL + fmt.Sprintf("?callback=RumAir&radius=%v&latitude=%f&longitude=%f", radius, lat, lon)

	return doAPIGet(citiesNearbyURL)
}

func getStationLocation(station *AirStation) (latitude string, longitude string, err error) {
	type LimitedOneValueResponse struct {
		End    int `json:"end"`
		Start  int `json:"start"`
		Values [][]struct {
			V string `json:"v"`
		} `json:"values"`
		Vars []string `json:"vars"`
	}

	resp := &LimitedOneValueResponse{}

	//https://pmpro.dacsystem.pl/webapp/data/averages?type=chart&start=1561939200&end=1561949200&vars=27LON
	curr, currMinus2h := nowAndMInus2hInUnixTimestamp()
	stationLatitudeURI := pmproSystemBaseApiURL + fmt.Sprintf("/averages?type=chart&start=%s&end=%s&vars=%s", curr, currMinus2h, station.LatitudeSensor)

	bytesRead, err := doAPIGet(stationLatitudeURI)

	if err != nil {
		return "", "", err
	}

	err = json.Unmarshal(bytesRead, resp)
	if err != nil {
		return "", "", err
	}

	latitude = resp.Values[0][0].V

	stationLongitudeURI := pmproSystemBaseApiURL + fmt.Sprintf("/averages?type=chart&start=%s&end=%s&vars=%s", curr, currMinus2h, station.LongitudeSensor)

	bytesRead, err = doAPIGet(stationLongitudeURI)

	if err != nil {
		return "", "", err
	}

	err = json.Unmarshal(bytesRead, resp)
	if err != nil {
		return "", "", err
	}

	longitude = resp.Values[0][0].V

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

func nowAndMInus2hInUnixTimestamp() (current string, currentMinus2h string) {
	current = string(time.Now().Unix())
	currentMinus2h = string(time.Now().Add(-time.Hour).Unix())

	return
}
