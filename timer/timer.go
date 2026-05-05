package timer

import (
	"fmt"
	"io"
	"sync"
	"time"
)

const (
	//timer clock states
	PAUSE = "PAUSE"
	RESUME = "RESUME"
	STOP = "STOP"

	//Interval states
	WORK = "WORK"
	SHORT_REST = "SHORT_REST"
	LONG_REST = "LONG_REST"

)


type Timer interface {
 NewTimer(config config)
// 	Start(io.Writer, chan string)
// 	Resume(io.Writer, chan string)
// 	Pause(io.Writer, chan string)
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
	p.running = false 
	p.completed = false
	control <- RESUME
}

//Stop: resets the timer state. This is from an interrupt that the user passes 
// so it does not necessarily signal the completion of an interval.
func (p *PomoTimer) Stop(out io.Writer, control chan string){
	p.running = false
	p.completed = false
	control <- STOP
}


// Core to what makes the timer tick. This helper will  
// tick the clock and reduce the currentTime until the time is done.
// the control channel can send events for play / pause / resume / stop
func (p *PomoTimer) Time(control chan string, out io.Writer, wg *sync.WaitGroup){
	defer wg.Done()
	ticker := time.NewTicker(p.tickSpeed)

	for{
		select{
		case <- ticker.C://signals per 'tick' of the timer
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
		case action := <-control:// signals when an action is provided.
			switch action {
			case PAUSE:
				p.running = false
				fmt.Println("Paused")
			case RESUME:
				p.running = true
				fmt.Println("Resumed")
			case STOP: 
				ticker.Stop()
				return
			}

		}
	}
}
