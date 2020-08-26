package tbox

import "time"

// get from now to tomorrow midnight during
func FromNowToTomorrowMidnight() time.Duration {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	tomorrow := today.AddDate(0, 0, 1)
	return tomorrow.Sub(now)
}
