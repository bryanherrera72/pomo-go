package timer

import "time"

type config struct {
	WorkDuration     time.Duration //configured time for a work duration. Default 25 minutes
	RestDuration     time.Duration //configured time for a short rest duration. Default 5 minutes
	LongRestDuration time.Duration //configured time for a long rest duration. Default 15 minutes
	Pomos            int           // total number of cycles. Default 4.
	GoalPomos        int           // desired number of pomodorros. Default 4.
}

//NewConfig() returns a new configuration for the pomodorro timer. With defaults set to the generic pomodorro defaults.
func NewConfig() config {
	config := config{}

	config.WorkDuration = 25 * time.Minute
	config.RestDuration = 5 * time.Minute
	config.LongRestDuration = 15 * time.Minute
	config.Pomos = 4
	config.GoalPomos = 4

	return config
}
