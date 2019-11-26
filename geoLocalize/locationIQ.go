package geolocalize

import (
	"fmt"
)

var locationiqBaseApiURL string = "https://locationiq.org/v1/reverse.php?key=e281731b38bb74"

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

/*func getCitiesNearby(lat float64, lon float64) (citiesNearby []string, err error) {
	radius := 0
	var reverseGeocodingStringedResponce string
	var bytesRead []byte

	for until := true; until; until = (len(reverseGeocodingStringedResponce)) < 5 {
	}
}*/
