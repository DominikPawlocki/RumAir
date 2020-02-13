package geolocalize

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dompaw/RumAir/airStations"
)

var locationiqBaseApiURL string = "http://locationiq.org/v1/reverse.php?key=e281731b38bb74"

func LocalizeStationsLocIQ(stations map[string]*airStations.AirStation) (result map[string]*LocalizedAirStation, err error) {
	result = map[string]*LocalizedAirStation{}
	for id, station := range stations {
		if station.LatitudeSensor != "" && station.LongitudeSensor != "" {
			//LocationIQ free API has to be called in at min 1 sec interval, it returns 400 if not
			SleepWithOutputDotsOnConsole(1 * time.Second)
			if localizedStation, err := LocalizeStationLocIQ(station); err == nil {
				result[id] = localizedStation
			}
		}
	}
	return
}

func LocalizeStationLocIQ(station *airStations.AirStation) (result *LocalizedAirStation, err error) {
	result = &LocalizedAirStation{Station: station}
	if result.Lat, result.Lon, err = getStationCoordinates(station); err == nil && result.Lat != 0 && result.Lon != 0 {
		result.CitiesNearby, err = getCitiesNearbyLocIQ(result.Lat, result.Lon)
		return
	}
	return result, fmt.Errorf("Can't localize station %v lat and lon. \n", station.ID)
}

func getReverseGeocodedCitiesLocIQ(lat float64, lon float64) (bytesRead []byte, err error) {
	//https://locationiq.org/v1/reverse.php?key=e281731b38bb74&lat=54.5987&lon=18.26669336&format=json
	locationIqURL := locationiqBaseApiURL + fmt.Sprintf("&lat=%f&lon=%f&format=json", lat, lon)
	fmt.Printf("ENDPOINT : %s. \n", locationIqURL)

	return doAPIGet(locationIqURL)
}

func getCitiesNearbyLocIQ(lat float64, lon float64) (citiesNearby []string, err error) {
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

	var bytesRead []byte
	var result *LocationIQReverseGeoResponse
	bytesRead, err = getReverseGeocodedCitiesLocIQ(lat, lon)

	if err != nil {
		fmt.Printf("Error during ReadAll bytesRead: %s err: %v. \n", bytesRead, err)
	}

	if len(bytesRead) > 0 {
		err = json.Unmarshal(bytesRead, &result)
		if err != nil {
			fmt.Printf("Error during deserializing occured. Data from `../table=Measurement&v=2`. Error is : %v", err)
			return
		}

		cityOrTown := result.Address.Town
		if cityOrTown == "" {
			cityOrTown = result.Address.County
		}

		citiesNearby = append(citiesNearby, fmt.Sprintf("%s, %s", cityOrTown, result.Address.Suburb))
	}
	return
}
