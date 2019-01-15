package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type DefaultSleeper struct{}

type Sleeper interface {
	Sleep()
}

func (ds *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i >= 1; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalWord)
	sleeper.Sleep()
}

func main() {
	defaultSleeper := &DefaultSleeper{}
	Countdown(os.Stdout, defaultSleeper)
}
