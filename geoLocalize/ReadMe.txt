The 'pmpro.dacsystem.pl' system has air monitoring stations in several cities all over a Poland.
Cities which uses this system usually has webpages exposed like 'rumia.powietrze.eu', monitoring current air status with short (3 days) history.
But there is no list of a pages like that, and a call 'https://pmpro.dacsystem.pl/webapp/json/do?table=Measurement&v=2' reveals there are like 60 stations over the Poland.
The question is : which stations are where ?



This code answers this question, with usage of two public geocoding APIs.

1)  from  Geobytes, entry : geoBytes.go LocalizeStationsGeoBytes
Their API is mostly used for geolocating IP adresses probably, but also has a possibility to find nearest city by latitutude/longitude .
Its simply and just finds biggest cities nearby, no more details.

2) from locationIQ, entry : locationIQ.go LocalizeStationsLocIQ
This API gives details about lon/lat provided like town, district, street, even nr of the building ...

You can choose which one to use.
Remember, most of the stations areny localizable (doesn't expose lat/lon)


Thats it ! Now I know which stations nearby my place Im interrested in grabbing history from !
