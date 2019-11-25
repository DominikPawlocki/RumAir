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

package geolocalize

import (
	"fmt"
)

var locationiqBaseApiURL string = "https://locationiq.org/v1/reverse.php?key=e281731b38bb74"

var pmproSystemBaseApiURL string = "https://pmpro.dacsystem.pl/webapp/data"

func getReverseGeocodedCitiesLocationIQ(radius int, lat float64, lon float64) (bytesRead []byte, err error) {
	type LocationIQReverseGeoResponse struct {
		Address struct {
			CityDistrict string `json:"city_district"`
			Country      string `json:"country"`
			CountryCode  string `json:"country_code"`
			County       string `json:"county"`
			HouseNumber  string `json:"house_number"`
			Postcode     string `json:"postcode"`
			Road         string `json:"road"`
			State        string `json:"state"`
			Suburb       string `json:"suburb"`
			Town         string `json:"town"`
		} `json:"address"`
		DisplayName string `json:"display_name"`
		Lat         string `json:"lat"`
		Licence     string `json:"licence"`
		Lon         string `json:"lon"`
	}

	//https://locationiq.org/v1/reverse.php?key=e281731b38bb74&lat=54.5987&lon=18.26669336&format=json
	locationIqURL := locationiqBaseApiURL + fmt.Sprintf("?callback=RumAir&radius=%v&latitude=%f&longitude=%f", radius, lat, lon)
	fmt.Printf("ENDPOINT : %s. \n", locationIqURL)

	return doAPIGet(locationIqURL)
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
