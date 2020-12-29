package airStations

import (
	"errors"
	"fmt"
	"sort"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

//remember that TestMain(m *testing.M) is also called for that module !
// MOCK GENERATING :
// mockgen.exe -destination=.\api\mocks_test1.go -package= api github.com/dompaw/Abax/api IStationsCapabiltiesFetcher
// or reflect mode and then copy paste from console : C:\.....\RumAir_Pmpro_Sensors_API> mockgen.exe github.com/dompaw/RumAir_Pmpro_Sensors_API/airStations IHttpAbstracter
func Test_Given_ErrorResponseFromDoAllMeasurmentsAPIcall_When_GetAllStationsCapabilities_Then_ResponseIsNilAndError(t *testing.T) {
	exampleMockErrorText := "timeout expired"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockIHttpAbstracter(ctrl)

	// Mock setting up
	m.
		EXPECT().
		DoHttpGetCall(allStationsMeasurmentsURL).
		Return(nil, errors.New(exampleMockErrorText)).
		AnyTimes()

	stations, err = GetAllStationsCapabilities(m)

	assert.Nil(t, stations)
	assert.Equal(t, exampleMockErrorText, err.Error(), fmt.Sprintf("Expected error like %s,but got %s in result", exampleMockErrorText, err.Error()))
}

func Test_GivenTwoStations_When_ShowStationsSensorsCodes_Then_AnswerIsCorrect(t *testing.T) {
	var localizableStation = &AirStation{
		ID: 2,
		Sensors: []SensorSimplified{
			SensorSimplified{
				ID:   1760,
				Code: "Sensor02_AA"},
			SensorSimplified{
				ID:   1760,
				Code: "Sensor02_BB"}},
		SensorsCount: 2}

	var notLocalizableStation = &AirStation{
		ID: 21,
		Sensors: []SensorSimplified{
			SensorSimplified{
				ID:   1760,
				Code: "HES21_PA"}},
		SensorsCount: 1}

	var stations = map[string]*AirStation{"02": localizableStation, "021": notLocalizableStation}

	actual := ShowStationsSensorsCodes(stations)

	var expectedFirst = "Station : 02 can :  Sensor02_AA Sensor02_BB"
	var expectedSecond = "Station : 021 can :  HES21_PA"

	assert.Equal(t, expectedFirst, actual[0], fmt.Sprintf("There should %s and is %s", expectedFirst, actual[0]))
	assert.Equal(t, expectedSecond, actual[1], fmt.Sprintf("There should %s and is %s", expectedSecond, actual[1]))

	assert.Condition(t, func() (success bool) { return sort.StringsAreSorted(actual) }, fmt.Sprint("Should be sorted."))
}

func Test_Given_StationNumber04_And_ErrorResponseFromDoAllMeasurmentsAPIcall_When_GetStationSensors_Then_ResponseIsNilAndError(t *testing.T) {
	exampleMockErrorText := "timeout expired"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockIHttpAbstracter(ctrl)

	// Mock setting up
	m.
		EXPECT().
		DoHttpGetCall(allStationsMeasurmentsURL).
		Return(nil, errors.New(exampleMockErrorText)).
		AnyTimes()

	actual, err := GetStationSensors(m, "04")

	assert.Nil(t, actual)
	assert.NotNil(t, err)
	assert.Equal(t, exampleMockErrorText, err.Error(), fmt.Sprintf("Expected error like %s,but got %s in result", exampleMockErrorText, err.Error()))
}

func Test_Given_StationNumber04_And_Sensor041_DoesSensorBelongsToStation_ReturnsFalse(t *testing.T) {
	var smt = Sensor{Code: "01"}
	result := doesSensorBelongsToStation(smt, "04")

	assert.False(t, result)
}

func Test_Given_StationNumber041_And_Sensor04_DoesSensorBelongsToStation_ReturnsFalse(t *testing.T) {
	var smt = Sensor{Code: "01"}
	result := doesSensorBelongsToStation(smt, "04")

	assert.False(t, result)
}

func Test_Given_StationNumber07_And_Sensor07_DoesSensorBelongsToStation_ReturnsTrue(t *testing.T) {
	var smt = Sensor{Code: "07"}
	result := doesSensorBelongsToStation(smt, "07")

	assert.True(t, result)
}

func Test_Given_StationNumber07_And_SensorWithLettersOnly_DoesSensorBelongsToStation_ReturnsFalse(t *testing.T) {
	var smt = Sensor{Code: "blaBla"}
	result := doesSensorBelongsToStation(smt, "07")

	assert.False(t, result)
}
