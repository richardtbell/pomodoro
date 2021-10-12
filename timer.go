package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"

	"github.com/gen2brain/beeep"
)

type timer struct {
	start        time.Time
	inWorkMode   bool
	workDuration int
	restDuration int
}

func (t timer) getElapsedTimeInSeconds() int {
	return int(time.Since(t.start).Seconds())
}

func (t *timer) switchMode() {
	t.start = time.Now()
	t.inWorkMode = !t.inWorkMode
}

func (t timer) alert() {
	message := "Take a break"
	if !t.inWorkMode {
		message = "Back to work"
	}
	fmt.Println(message)
	os := runtime.GOOS
	if os == "darwin" && !isQuietTime() {
		go exec.Command("say", message).Output()
	}
	beeep.Notify(message, message, "assets/information.png")
}

func (t timer) shouldSwitchMode(elapsed int) bool {
	return elapsed == t.getDuration()
}

func (t timer) getDuration() int {
	duration := t.workDuration
	if !t.inWorkMode {
		duration = t.restDuration
	}
	return duration
}
func (t timer) getMode() string {
	mode := "Work"
	if !t.inWorkMode {
		mode = "Rest"
	}
	return mode
}

func (t timer) printTimeRemaining(elapsed int) {
	timeRemaining := t.getDuration() - elapsed
	minutes := timeRemaining / 60
	seconds := timeRemaining - minutes*60
	fmt.Printf("\r%v: %02d:%02d", t.getMode(), minutes, seconds)
}
