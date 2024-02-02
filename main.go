package main

import (
	"fmt"
	"os"

	"github.com/ashiqYousuf/go-with-tests/mocking"
)

func main() {
	sleeper := &mocking.DefaultSleeper{}
	// sleeper := &mocking.ConfigurableSleeper{time.Second * 1, time.Sleep}
	mocking.Countdown(os.Stdout, sleeper)
	fmt.Println()
}
