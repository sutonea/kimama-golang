package timecalc

import (
	"time"
)

// strDate is base.
// days : add days
func AddDay(strDate string, days int) time.Time {
	baseTime, _ := time.Parse(time.RFC3339, strDate)
	return baseTime.AddDate(0, 0, days)
}
