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

package sensors

import (
	"errors"
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

var cronInstance *cron.Cron = cron.New(cron.WithSeconds())

/*func AddSensorToCron(c *cron.Cron, sensorId strings, offsetInSeconds int) cron.EntryID {
	fmt.Println("---------------CRON ----------------------")
	id, err := c.AddFunc("30 * * * *", func() { fmt.Println("Every hour on the half hour") })
	if err != nil {
		fmt.Printf("Error occured when adding sensor %v to CRON !", sensorId)
		panic(err)
	}

	fmt.Printf("Sensor %v added to Cron.", sensorId)

	return id
}*/

func GetCron() *cron.Cron {
	return cronInstance
}

func AddSensorToCron(sensorId string, offsetInSeconds int) (cron.EntryID, error) {
	timeOffset := 6 * offsetInSeconds
	id, err := cronInstance.AddFunc("10 * * * * *", func() { fmt.Printf("Time %v \n", time.Now()) })
	if err != nil {
		fmt.Printf("Error occured when adding sensor %v to CRON !", sensorId)
		panic(err)
	}
	return id, nil
}

func StartCron() (int, error) {
	numberOfEntriesInCron := len(cronInstance.Entries())
	if numberOfEntriesInCron > 0 {
		cronInstance.Start()
		return numberOfEntriesInCron, nil
	}
	return -1, errors.New("Cron Not started !!")
}
