package examples

import "time"

func CurrentDate() string {
	now := time.Now()
	return now.Format("January 2, 2006")
}
