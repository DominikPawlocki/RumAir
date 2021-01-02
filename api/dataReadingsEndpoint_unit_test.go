package api

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_dayMonthYearParser(t *testing.T) {
	tables := []struct {
		day   string
		month string
		year  string
	}{
		{"", "aa", "2020"},
		{"-1", "0", "3000"},
		{"2", "13", "2020"},
		{"15", "05", "1070"},
	}

	for _, table := range tables {
		_, err := dayMonthYearParser(table.day, table.month, table.year)
		assert.Error(t, err)
	}
}

func Test_GivenIncorrectSensorCodesQueryStr_WhenFilteringIt_ReturnsEmptyString(t *testing.T) {
	var incorrectQueryStrings = []string{
		"",
		"&bla bla blaBla",
		"56s84 f7 3678&8763io ",
		" mnbzxc51and&42",
	}

	for _, qryStr := range incorrectQueryStrings {
		actual := filterQueryStringToGetSensorsCodesBelongingToStationOnly(qryStr, "05")
		assert.Nil(t, actual, fmt.Sprintf("There should be nil and is not."))
	}
}

func Test_GivenOKSensorCodesQueryStr_WhenFilteringIt_ReturnsCorrectResult(t *testing.T) {
	var s1 = "05Sensor1"
	var s2 = "05aaaensor2"
	var s3 = "04_Saensor3_doesntBelongTo05"
	var station05SensorCodesOnly = fmt.Sprintf("%s,%s,%s", s1, s2, s3)

	actual := filterQueryStringToGetSensorsCodesBelongingToStationOnly(station05SensorCodesOnly, "05")
	assert.NotEmpty(t, actual)
	assert.Equal(t, []string{s1, s2}, actual, fmt.Sprintf("There should be nil and is not."))
}

func Test_GivenDate_WhenConvertingToEpoch_ReturnsCorrectDate(t *testing.T) {
	var testTime = time.Date(2020, 05, 18, 05, 40, 15, 0, time.UTC)

	beginOfDay, endOfDay := getBeginAndEndofTheDayInUnixEpoch(testTime)
	assert.NotEmpty(t, beginOfDay)
	assert.NotEmpty(t, endOfDay)
	assert.Equal(t, time.Date(testTime.Year(), testTime.Month(), testTime.Day(), 0, 0, 0, 0, time.UTC).Unix(), beginOfDay)
	assert.Equal(t, time.Date(testTime.Year(), testTime.Month(), testTime.Day(), 23, 59, 59, 0, time.UTC).Unix(), endOfDay)
}
