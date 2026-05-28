package main

import (
	"fmt"
	"os"

	"charm.land/bubbles/timer"
	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
)

type timermodel struct {
	// timer    timer.Model
	keymap   keymap
	help     help.Model
	quitting bool
}
type keymap struct{
	start key.Binding
	stop key.Binding
	pause key.Binding
	quit key.Binding
}

func (m timermodel) Init() tea.Cmd{
	return nil
}

//TODO: need this to react to ticks probably
func (m timermodel) Update(msg tea.Msg) (tea.Model, tea.Cmd){
	switch msg := msg.(type){
	case timer.TickMsg:
			 
	}
	return nil, nil
}
func(m timermodel) View(){
}

func main() {
	// config := timer.NewTestConfig()
	// config.WorkDuration = 500 * time.Millisecond
	// timer := timer.NewTimer(config)

	// //this is how the timer will behave. we'll need a wait group for the timer. 
	// //and its controls
	// control := make(chan string)
	// defer close(control)

	// var wg sync.WaitGroup
	// wg.Add(1)
	// go timer.Time(control, os.Stdout, &wg)
	// timer.Start(os.Stdout, control)
	// wg.Wait()
	m := timermodel{

	}

	if _, err := tea.NewProgram(m).Run(); err != nil{
		fmt.Println("An Error occurred. Exiting with error: ", err)
		os.Exit(1)
	}
}