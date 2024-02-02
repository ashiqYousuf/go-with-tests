package mocking

import (
	"fmt"
	"io"
	"time"
)

// ? Let's define our dependency as an interface.
// ? This lets us then use a real Sleeper in main and a spy sleeper in our tests.
type Sleeper interface {
	Sleep()
}

// ? Now we need to make a mock of it for our tests to use.
// type SpySleeper struct {
// 	Calls int
// }

// func (s *SpySleeper) Sleep() {
// 	s.Calls++
// }

// ? Let's create a real sleeper which implements the interface we need

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(time.Second * 1)
}

/*`
Countdown should sleep before each next print, e.g:
Print N
Sleep
Print N-1
Sleep
Print Go!`
*/

const write = "write"
const sleep = "sleep"

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

// !Let's first create a new type for ConfigurableSleeper that accepts what
// !we need for configuration and testing

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

const finalWord = "Go!"
const countdownStart = 3

// ? This Countdown when used in main, writes to stdout and when
// ? used in TestCountdown writes to buffer (some physical store)
func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i >= 1; i-- {
		fmt.Fprintln(writer, i)
		// ! Slow tests ruin developer productivity.
		// ! Are we happy with 3s added to the test run for every new test of Countdown?
		// ! If we can mock time.Sleep we can use dependency injection to use it instead of a "real" time.Sleep
		// ! and then we can spy on the calls to make assertions on them.
		sleeper.Sleep()
	}
	fmt.Fprintf(writer, finalWord)
}
