package sensors

//https://pmpro.dacsystem.pl/webapp/data/averages?type=chart&avg=2h&start=1561939200&end=1561949200&vars=38LAT

/* data from air sensors collected by https://pmpro.dacsystem.pl/ are organized as : stationId+dataId.
The call: https://pmpro.dacsystem.pl/webapp/json/do?table=Measurement&v=2  gives all data types collected by sensors.
But which sensor is where ? There a LAT and LON data helps.
Getting this data per sensor ids 01 to 40 (highest sensorId on 11/2019) gives us locations !*/

/*STACJE GDYNIA :
"vars":

[
"004CO:A1h", Referencyjna, AM
"004SO2:A1h",
"004NO2:A1h",
"004O3:A1h",
"009PA:A1h", referencyjna
"009TEMP:A1h",
"009WD:V1h",
"009WS:V1h",
"009NO2:A1h",
"009O3:A1h",
"009SO2:A1h",
"009PM10:A1h",
"20PM10A_W_k:A1h", Obłuże, Krawiecka 35
"20PM25A_W_k:A1h",
"21PM10A_W_k:A1h",
"22PM10A_W_k:A1h",
"22PM25A_W_k:A1h",
"23PM10A_W_k:A1h",
"23PM25A_W_k:A1h",
"24PM10A_W_k:A1h",
"24PM25A_W_k:A1h",
"25PM10A_W_k:A1h",
"25PM25A_W_k:A1h",
"26PM10A_W_k:A1h",
"26PM25A_W_k:A1h",
"27PM10A_W_k:A1h",
"27PM25A_W_k:A1h",
"28PM10A_W_k:A1h",
"28PM25A_W_k:A1h",
"29PM10A_W_k:A1h",
"29PM25A_W_k:A1h",
"004PM10:A1h",
"41PM10A_W_k:A1h",
"41PM25_W_k:A1h",
"42PM10A_W_k:A1h",
"42PM25A_W_k:A1h",
"43PM10A_W_k:A1h",
"43PM25A_W_k:A1h",
"21PM25A_W_k:A1h"
]*/

/*STACJE RUMIA :
"vars":

[
"04HUMID_O:A1h",
"04PRESS_O:A1h",
"05HUMID_O:A1h",
"05PRESS_O:A1h",
"06HUMID_O:A1h",
"06PRESS_O:A1h",
"07HUMID_O:A1h",
"07PRESS_O:A1h",
"08HUMID_O:A1h",
"08PRESS_O:A1h",
"05PM10A_6_k:A1h",
"06PM10A_6_k:A1h",
"07PM10A_6_k:A1h",
"08PM10A_6_k:A1h",
"04PM10A_6_k:A1h",
"009PA:A1h",
"009TEMP:A1h",
"009WD:V1h",
"009WS:V1h",
"04PM25A_6_k:A1h",
"05PM25A_6_k:A1h",
"06PM25A_6_k:A1h",
"07PM25A_6_k:A1h",
"08PM25A_6_k:A1h"

dodatkowo jest :
06TEMP_O itd....

]*/

/* STACJE WEJHEROWO :
vars":

[
"30HUMID_O:A1h",
"30PRESS_O:A1h",
"31HUMID_O:A1h",
"31PRESS_O:A1h",
"32HUMID_O:A1h",
"32PRESS_O:A1h",
"33HUMID_O:A1h",
"33PRESS_O:A1h",
"35HUMID_O:A1h",
"35PRESS_O:A1h",
"36HUMID_O:A1h",
"36PRESS_O:A1h",
"009PA:A1h",
"009TEMP:A1h",
"009WD:V1h",
"009WS:V1h",
"30PM10A_W_k:A1h",
"30PM25A_W_k:A1h",
"31PM10A_W_k:A1h",
"31PM25A_W_k:A1h",
"32PM10A_W_k:A1h",
"32PM25A_W_k:A1h",
"33PM10A_W_k:A1h",
"33PM25A_W_k:A1h",
"35PM10A_W_k:A1h",
"35PM25A_W_k:A1h",
"36PM10A_W_k:A1h",
"36PM25A_W_k:A1h"
]*/
