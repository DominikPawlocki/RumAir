package airStations

import (
	"errors"
	"fmt"
	"sort"
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

//remember that TestMain(m *testing.M) is also called for that module !
// MOCK GENERATING :
// mockgen.exe -destination=.\api\mocks_test1.go -package= api github.com/dompaw/RumAir_Pmpro_Sensors_API/airStations IHttpAbstracter
// or reflect mode and then copy paste from console : C:\.....\RumAir_Pmpro_Sensors_API> mockgen.exe github.com/dompaw/RumAir_Pmpro_Sensors_API/airStations IHttpAbstracter
func Test_Given_ErrorResponseFromDoAllMeasurmentsAPIcall_When_GetAllStationsCapabilities_Then_ResponseIsNilAndError(t *testing.T) {
	exampleMockErrorText := "timeout expired"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockIHttpAbstracter(ctrl)

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
	var sensorCode1_Station2 = "Sensor02_AA"
	var sensorCode2_Station2 = "Sensor02_BB"
	var sensorCode1_Station21 = "HES21_PA"

	var localizableStation = &AirStation{
		ID: 2,
		Sensors: []SensorSimplified{
			SensorSimplified{
				ID:   1760,
				Code: sensorCode1_Station2},
			SensorSimplified{
				ID:   1760,
				Code: sensorCode2_Station2}},
		SensorsCount: 2}

	var notLocalizableStation = &AirStation{
		ID: 21,
		Sensors: []SensorSimplified{
			SensorSimplified{
				ID:   1760,
				Code: sensorCode1_Station21}},
		SensorsCount: 1}

	var stations = map[string]*AirStation{"02": localizableStation, "021": notLocalizableStation}

	actual := ShowStationsSensorsCodes(stations)

	var expectedFirst = fmt.Sprintf("Station : 02 can :  %s %s", sensorCode1_Station2, sensorCode2_Station2)
	var expectedSecond = fmt.Sprintf("Station : 021 can :  %s", sensorCode1_Station21)

	assert.Equal(t, expectedFirst, actual[0], fmt.Sprintf("There should %s and is %s", expectedFirst, actual[0]))
	assert.Equal(t, expectedSecond, actual[1], fmt.Sprintf("There should %s and is %s", expectedSecond, actual[1]))

	assert.Condition(t, func() (success bool) { return sort.StringsAreSorted(actual) }, fmt.Sprint("Should be sorted."))
}

func Test_GivenTwoSensors_When_GetStationSensorCodesOnly_Then_AnswerIsCorrect(t *testing.T) {
	var sensorCode1 = "07_Test1"
	var sensorCode2 = "07_Test2"

	var mockedResponse = []byte(fmt.Sprintf(`{"success":true,"totalCount":2,"message":"","data": [ {"code":"%s"},{"code":"%s"}]}`, sensorCode1, sensorCode2))

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockIHttpAbstracter(ctrl)

	m.
		EXPECT().
		DoHttpGetCall(allStationsMeasurmentsURL).
		Return(mockedResponse, nil).
		AnyTimes()

	actual, _ := GetStationSensorCodesOnly(m, "07")

	assert.NotNil(t, actual)
	assert.Equal(t, 2, len(actual), fmt.Sprintf("There should len of 2 and is %v", len(actual)))
	assert.Equal(t, sensorCode1, actual[0], fmt.Sprintf("There should be %s and is %s", sensorCode1, actual[0]))
	assert.Equal(t, sensorCode2, actual[1], fmt.Sprintf("There should be %s and is %s", sensorCode2, actual[1]))
}

func Test_GivenTwoSensorsWithSameStartDate_When_GetSensorStartTimeAndCode_Then_AnswerIsCorrect(t *testing.T) {
	var expectedUnixTime int64 = 1609358600
	var sensorCode1 = "07_Test1"
	var sensorCode2 = "07_Test2"

	var mockedResponse = []byte(
		fmt.Sprintf(`{"success":true,"totalCount":2,"message":"","data": [ {"code":"%s", "start_date":%d},{"code":"%s", "start_date":%d}]}`,
			sensorCode1, expectedUnixTime, sensorCode2, expectedUnixTime))

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockIHttpAbstracter(ctrl)

	m.
		EXPECT().
		DoHttpGetCall(allStationsMeasurmentsURL).
		Return(mockedResponse, nil).
		AnyTimes()

	actual, _ := GetSensorStartTimeAndCode(m, "07")

	assert.NotNil(t, actual)
	assert.Equal(t, time.Date(2020, 12, 30, 21, 03, 20, 0, time.UTC), actual[expectedUnixTime].StartDate,
		fmt.Sprintf("There should %v and is %s", expectedUnixTime, actual[expectedUnixTime].StartDate))
	assert.Equal(t, 2, len(actual[expectedUnixTime].Codes), fmt.Sprintf("There should len of 2 and is %v", len(actual[expectedUnixTime].Codes)))
	assert.Equal(t, sensorCode1, actual[expectedUnixTime].Codes[0], fmt.Sprintf("There should be %s and is %s", sensorCode1, actual[expectedUnixTime].Codes[0]))
	assert.Equal(t, sensorCode2, actual[expectedUnixTime].Codes[1], fmt.Sprintf("There should be %s and is %s", sensorCode2, actual[expectedUnixTime].Codes[1]))
}

func Test_Given_StationNumber04_And_ErrorResponseFromDoHttpGetCall_When_GetStationSensors_Then_ResponseIsNilAndError(t *testing.T) {
	exampleMockErrorText := "timeout expired"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockIHttpAbstracter(ctrl)

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
