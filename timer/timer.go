package timer

import (
	"fmt"
	"io"
	"sync"
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
	currentTime 	 time.Duration
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
func (p *PomoTimer) Start(out io.Writer, control chan string) {
	p.currentTime = p.WorkDuration
	control <- "RESUME"
}


// Core to what makes the timer tick. This helper will  
// tick the clock and reduce the currentTime until the time is done.
// the control channel can send events for play / pause / resume / stop
func ( p *PomoTimer)Time(control chan string, out io.Writer, wg *sync.WaitGroup){
	defer wg.Done()
	ticker := time.NewTicker(p.tickSpeed)

	for{
		select{
		case <- ticker.C:
			if(p.running){
				timeAsBytes := []byte(p.currentTime.String())
				out.Write(timeAsBytes)
				p.currentTime -= p.tickSpeed
				if p.currentTime <= 0{
					p.running = false
					p.completed = true
					ticker.Stop()
					return
				}
			}
		case action := <-control:
			if action == "PAUSE"{
				p.running = false
				fmt.Println("Paused")
			} else if action == "RESUME"{
				p.running = true
				fmt.Println("Resumed")
			} else if action == "STOP"{
				return
			} 
		}
	}
}
