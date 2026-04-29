package timer

import "time"

const standardTimeMeasure = 1 * time.Minute
const testTimeMeasure = 1 * time.Millisecond

const standardTickInterval = 1 * time.Second
const testTickInterval = 1 * time.Millisecond

type config struct {
	WorkDuration     time.Duration //configured time for a work duration. Default 25 minutes
	RestDuration     time.Duration //configured time for a short rest duration. Default 5 minutes
	LongRestDuration time.Duration //configured time for a long rest duration. Default 15 minutes
	Pomos            int           // total number of cycles. Default 4.
	GoalPomos        int           // desired number of pomodorros. Default 4.
	tickSpeed 		 time.Duration // speed of the ticker interval. default goes by second. for tests, the interval is milliseconds. 
}

//NewConfig() returns a new configuration for the pomodorro timer. With defaults set to the generic pomodorro defaults.
func NewConfig() config {
	config := config{}

	config.WorkDuration = 25 * standardTimeMeasure
	config.RestDuration = 5 * standardTimeMeasure
	config.LongRestDuration = 15 * standardTimeMeasure
	config.Pomos = 4
	config.GoalPomos = 4
	config.tickSpeed = standardTickInterval
	return config
}

func NewTestConfig() config {
	config := config{}

	config.WorkDuration = 25 * testTimeMeasure
	config.RestDuration = 5 * testTimeMeasure
	config.LongRestDuration = 15 * testTimeMeasure
	config.Pomos = 4
	config.GoalPomos = 4
	config.tickSpeed = testTickInterval
	return config
}
