package tsync

import (
	"sync"
)

/*
? A Mutex is a mutual exclusion lock.
? The zero value for a Mutex is an unlocked mutex
? A Mutex must not be copied after first use.

! go vet:- alert you to some subtle bugs in your code
*/

type Counter struct {
	mu    sync.Mutex
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
