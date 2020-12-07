package api

import (
	"testing"

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
