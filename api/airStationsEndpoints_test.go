package api

import (
	"os"
	"testing"
)

var stations map[string]*AirStation
var err error

func setup() {
	stations, err = GetAllStationsCapabilities()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	//shutdown()
	os.Exit(code)
}

func Test_ShowStationsSensorsCodesHandler(t *testing.T) {
	actual := ShowStationsSensorsCodesHandler()
}
