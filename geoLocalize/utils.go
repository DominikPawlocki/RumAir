package geolocalize

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

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

func nowAndMInus2hInUnixTimestamp() (current int64, currentMinus2h int64) {
	current = time.Now().Unix()
	currentMinus2h = time.Now().Add(-3 * time.Hour).Unix()

	return
}
