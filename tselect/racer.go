package tselect

import (
	"errors"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

// ? struct{} Empty struct is just used for signaling or synchronization purposes
// ? chan struct{} is the smallest data type available from a memory perspective.

// ! For channels the zero value is nil and if you try and send to it
// ! with <- it will block forever because you cannot send to nil channels

func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}

func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", errors.New("Server took too long to respond")
	}
}

func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}
