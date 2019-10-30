package server

import (	
	"github.com/robfig/cron/v3"
)

type sensorReading struct {
	Symbol string  // Stock symbol
	Volume int     // Number of shares
	Price  float64 // Trade price
	Buy    bool    // true if buy trade, false if sell trade
}

//dont !!!!!!!!!!!!!!!!
