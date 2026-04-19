package timer

import "time"

// type Timer interface {
// 	Start(io.Writer)
// 	Resume(io.Writer)
// 	Stop(io.Writer)
// }

type Timer struct {
	WorkDuration     time.Duration
	RestDuration     time.Duration
	LongRestDuration time.Duration
	Pomos            int
	GoalPomos        int
	tickSpeed        time.Duration
	running          bool
	completed        bool
}

func NewTimer(config config) Timer {
	timer := Timer{}

	timer.WorkDuration = config.WorkDuration
	timer.RestDuration = config.RestDuration
	timer.LongRestDuration = config.LongRestDuration
	timer.Pomos = config.Pomos
	timer.GoalPomos = config.GoalPomos
	timer.tickSpeed = 1 * time.Second
	timer.running = false
	timer.completed = false

	return timer
}

// Start begins the timer using the set work duration.
// This is the assumed "beginning" of the timer session ie. fresh timer
// func (p *PomoTimer) Start(out io.Writer) {
// 	workTime := p.WorkDuration
// 	ticker := time.NewTicker(1 * time.Second)

// 	complete := make(chan bool, 1)

// 	go func() {
// 		for {
// 			select {
// 			case <-complete:
// 				return
// 			case datetime := <-ticker.C:
// 				out.Write([]byte(datetime.String()))
// 				workTime -= time.Second
// 				if workTime <= 0 {
// 					complete <- true
// 				}
// 			}
// 		}
// 	}()
// 	<-complete
// 	ticker.Stop()
// }
