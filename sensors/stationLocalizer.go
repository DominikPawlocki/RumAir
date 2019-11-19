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
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type LocalizedAirStation struct {
	Station      *AirStation
	Lat          string
	Lon          string
	CitiesNearby []string
}

var geoBytesBaseApiURL string = "http://getnearbycities.geobytes.com/GetNearbyCities"
var geoBytesBaseApiURL string = "https://pmpro.dacsystem.pl/webapp/data"

func GetStationLonLat(station *AirStation) (result *LocalizedAirStation) {
	if station.LatitudeSensor && station.LongitudeSensor {
	

		if len(bytesRead) > 0 {
			err = json.Unmarshal(bytesRead, res *odpowiedzLATLON)
		
			if(err==null){
				station.
			}
		}
	}
}

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

func getStationLocation(station *AirStation) (latitude string, longitude string, err error)
{
	//https://pmpro.dacsystem.pl/webapp/data/averages?type=chart&start=1561939200&end=1561949200&vars=27LON
	stationLatitudeUri := geoBytesBaseApiURL + fmt.Sprintf("?callback=RumAir&radius=%v&latitude=%f&longitude=%f", radius, lat, lon)

	bytesRead, err = doAPIGet(stationLatitudeUri)

	if err != nil {
		return _,_, err
	}
	
	latitude = 

	stationLongitudeUri := geoBytesBaseApiURL + fmt.Sprintf("?callback=RumAir&radius=%v&latitude=%f&longitude=%f", radius, lat, lon)

	bytesRead, err = doAPIGet(stationLongitudeUri)

	if err != nil {
		return _,_, err
	}

	longitude = 

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

func ToUnixEpochTimestamp(){
	layout := "01/02/2006 3:04:05 PM"
    t, err := time.Parse(layout, "02/28/2016 9:03:46 PM")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(t.Unix())
}