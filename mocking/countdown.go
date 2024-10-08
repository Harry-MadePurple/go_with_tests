package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}
type DefaultSleeper struct{}

type ConfiguableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}
type SpyTime struct {
	durationSlept time.Duration
}

func (c *ConfiguableSleeper) Sleep() {
	c.sleep(c.duration)
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type SpyCountdownOperation struct {
	Calls []string
}

func (s *SpyCountdownOperation) Sleep() {
	s.Calls = append(s.Calls, sleep)
}
func (s *SpyCountdownOperation) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"
const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)

}

func main() {
	sleeper := &ConfiguableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
