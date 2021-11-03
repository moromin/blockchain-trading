package entity

import (
	"time"
)

// An interval represetns OHLC data for a particular Period and CloseTime.
type Interval struct {
	Period Period
	OHLC   OHLC

	// CloseTime is the time at which the Interval ended.
	CloseTime time.Time

	// VolumeBase is the amount of volume traded over this Interval, represented
	// in the base currency.
	VolumeBase float64

	// VolumeQuote is the amount of volume traded over this Interval, represented
	// in the quote currency.
	VolumeQuote float64
}

// Period is the number of seconds in an Interval.
type Period string

// The following constants are all available Period values for IntervalsUpdate.
const (
	Period1M         Period = "60"
	Period3M         Period = "180"
	Period5M         Period = "300"
	Period15M        Period = "900"
	Period30M        Period = "1800"
	Period1H         Period = "3600"
	Period2H         Period = "7200"
	Period4H         Period = "14400"
	Period6H         Period = "21600"
	Period12H        Period = "43200"
	Period1D         Period = "86400"
	Period3D         Period = "259200"
	Period1WThursday Period = "604800"
	Period1WMonday   Period = "604800_Monday"
)

func (p Period) Duration() time.Duration {
	return periodDurations[p]
}

var periodDurations = map[Period]time.Duration{
	"60":            60 * time.Second,
	"180":           180 * time.Second,
	"300":           300 * time.Second,
	"900":           900 * time.Second,
	"1800":          1800 * time.Second,
	"3600":          3600 * time.Second,
	"7200":          7200 * time.Second,
	"14400":         14400 * time.Second,
	"21600":         21600 * time.Second,
	"43200":         43200 * time.Second,
	"86400":         86400 * time.Second,
	"259200":        259200 * time.Second,
	"604800":        604800 * time.Second,
	"604800_Monday": 604800 * time.Second,
}

// PeriodNames contains human-readable names for Period.
// e.g. Period1M = 60 (seconds) = "1m".
var PeriodNames = map[Period]string{
	Period1M:         "1m",
	Period3M:         "3m",
	Period5M:         "5m",
	Period15M:        "15m",
	Period30M:        "30m",
	Period1H:         "1h",
	Period2H:         "2h",
	Period4H:         "4h",
	Period6H:         "6h",
	Period12H:        "12h",
	Period1D:         "1d",
	Period3D:         "3d",
	Period1WMonday:   "1w",
	Period1WThursday: "1w",
}

// OHLC contains the open, low, high, and close prices for a given Interval
type OHLC struct {
	Open  float64
	High  float64
	Low   float64
	Close float64
}
