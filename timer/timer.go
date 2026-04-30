package timer

import (
	"io"
	"time"
)

type Timer interface {
 NewTimer(config config)
// 	Start(io.Writer)
// 	Resume(io.Writer)
// 	Stop(io.Writer)
}

type PomoTimer struct {
	WorkDuration     time.Duration
	RestDuration     time.Duration
	LongRestDuration time.Duration
	Pomos            int
	GoalPomos        int
	tickSpeed        time.Duration
	running          bool
	completed        bool
}

//NewPomoTimer: returns an instance of our base pomodorro timer. 
func NewTimer(config config) PomoTimer {
	timer := PomoTimer{}

	timer.WorkDuration = config.WorkDuration
	timer.RestDuration = config.RestDuration
	timer.LongRestDuration = config.LongRestDuration
	timer.Pomos = config.Pomos
	timer.GoalPomos = config.GoalPomos
	timer.tickSpeed = config.tickSpeed
	timer.running = false
	timer.completed = false

	return timer
}

// Begins a fresh timer. Initializes state of the timer (running == completed == false)
// 
func (p *PomoTimer) Start(out io.Writer) {
	workTime := &p.WorkDuration
	ticker := time.NewTicker(p.tickSpeed)

	complete := make(chan bool, 1)
	go tickHelper(p, complete, out, workTime)
	<-complete
	ticker.Stop()
	p.completed = true
	p.running = false
}

// Core to what makes the timer tick. This helper will  
// tick the clock and reduce the currentTime until the time is done.
//
// p *PomoTimer: ref to the timer in use. 
//
// cmp chan bool: channel that receive a true bool when the timer is complete.
//
// out io.Writer: the output we wish to write to. Must implement Write() for io.Writer
//
// currentTime *time.Duration: Time that we wish to tick down. 
func tickHelper(p *PomoTimer, cmp chan bool, out io.Writer, currentTime *time.Duration){
	ticker := time.NewTicker(p.tickSpeed)
	for {
		select {
		case <-cmp:
			return
		case <-ticker.C:
			minutesAsBytes := []byte(currentTime.String())
			out.Write(minutesAsBytes)
			*currentTime -= p.tickSpeed
			if *currentTime <= 0 {
				cmp <- true
			}
		}
	}
	
}