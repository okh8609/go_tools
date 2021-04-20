package timer

import (
	"time"
)

func GetNowTime() time.Time {
	// location, _ := time.LoadLocation("Asia/Taipei")
	// return time.Now().In(location)
	return time.Now()
}

func GetCalcTime(currTime time.Time, dur string) (time.Time, error) {
	duration, err := time.ParseDuration(dur)
	if err != nil {
		return time.Time{}, err
	}
	return currTime.Add(duration), nil
}
