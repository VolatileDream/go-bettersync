package bettersync

import (
	"sync"
)

type Unlocker func()

type Mutex[Guarded any] struct {
	m sync.Mutex
	v Guarded
}

// New creates a bettersync.Mutex using the provided function to
// construct the value inside before making the mutex available.
func New[Guarded any](fn func() Guarded) *Mutex[Guarded] {
	return &Mutex[Guarded]{
		m: sync.Mutex{},
		v: fn(),
	}
}

// Lock locks the mutex. If already locked, blocks the goroutine
// until the mutex is available.
func (m *Mutex[Guarded]) Lock() (*Guarded, Unlocker) {
	m.m.Lock()
	return &m.v, m.unlock
}

func (m *Mutex[Guarded]) unlock() {
	m.m.Unlock()
}
