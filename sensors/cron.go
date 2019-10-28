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
	"fmt"

	"github.com/robfig/cron/v3"
)

func AddSensorToCron(c *cron.Cron, sensorId string, offsetInSeconds int) int {

	id, err := c.AddFunc("30 * * * *", func() { fmt.Println("Every hour on the half hour") })
	inspect(c.Entries())

	/*for _, entry := range c.Entries() {
	if id == entry.ID {
		return entry
	}*/

	return id.ID
}
