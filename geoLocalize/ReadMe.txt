The 'pmpro.dacsystem.pl' system has air monitoring stations in several cities all over a Poland.
Cities which uses this system usually has webpages exposed like 'rumia.powietrze.eu', monitoring current air status with short (3 days) history.
But there is no list of a pages like that, and a call 'http://pmpro.dacsystem.pl/webapp/json/do?table=Measurement&v=2' reveals there are like 60 stations over the Poland.
The question is : which stations are where ?



This code answers this question, with usage of two public geocoding APIs.

1)  from  Geobytes, with : .../stations/locate/geobytes
Their API is mostly used for geolocating IP adresses probably, but also has a possibility to find nearest city by latitutude/longitude .
Its simply and just finds biggest cities nearby, no more details.

2) from locationIQ, with : "/stations/locate/locationIQ", "/stations/{id}/locate/locationIQ", "/stations/locate/locationIQ/numbersPerCity"
This API gives details about lon/lat provided like town, district, street, even nr of the building ...

Remember, most of the stations arent localizable (doesn't expose lat/lon)
