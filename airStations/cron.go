/*
CRON Expression Format
A cron expression represents a set of times, using 5 space-separated fields.
Field name   | Mandatory? | Allowed values  | Allowed special characters
----------   | ---------- | --------------  | --------------------------
Minutes      | Yes        | 0-59            | * / , -
Hours        | Yes        | 0-23            | * / , -
Day of month | Yes        | 1-31            | * / , - ?
Month        | Yes        | 1-12 or JAN-DEC | * / , -
Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?*/

package airStations

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/robfig/cron/v3"
)

var cronInstance *cron.Cron = cron.New(cron.WithSeconds())

func AddStationsToCron(stations []Station) {
	for i, station := range stations {
		if station.CronHandler != nil {
			sensorID, err := addStationToCron(i, station.CronHandler)

			if err == nil {
				fmt.Printf("Sensor %v added to Cron.\n", sensorID)
			}
		}
	}
}

func GetCron() *cron.Cron {
	return cronInstance
}

func addStationToCron(offsetInSeconds int, sensorDataFetcher func()) (cron.EntryID, error) {
	timeOffset := 9 * offsetInSeconds
	id, err := cronInstance.AddFunc(cronFormatBuilder(timeOffset), sensorDataFetcher)
	if err != nil {
		fmt.Println("Error occured when adding sensor to CRON !")
		panic(err)
	}
	return id, nil
}

func cronFormatBuilder(secondsOffset int) string {
	return strconv.Itoa(secondsOffset) + " * * * * *"
}

func StartCron() (numberOfEntriesInCron int, err error) {
	numberOfEntriesInCron = len(cronInstance.Entries())
	if numberOfEntriesInCron > 0 {
		cronInstance.Start()
		return numberOfEntriesInCron, nil
	}
	return -1, errors.New("Cron not started")
}
