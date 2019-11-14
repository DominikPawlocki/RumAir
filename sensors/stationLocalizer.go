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

var geoBytesBaseApiURL string = "http://getnearbycities.geobytes.com/GetNearbyCities"

func getCitiesNearby(lat float32, lon float32) (citiesNearby []string) {
	radius := 40
	query := "?callback=?&radius=40&latitude=54.5708&longitude=18.3878"
	var geoBytesQueryUrl = fmt.Sprintf("%s", geoBytesApiURL, sensorID)  + 
}
