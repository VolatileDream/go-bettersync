package bettersync

import (
	"testing"
)

func Test_Mutex_initialization(t *testing.T) {
	m := &Mutex[int]{}

	g, unlock := m.Lock()
	defer unlock()

	if g == nil {
		t.Fatal("Got nil for guarded value")
	}
	if *g != 0 {
		t.Errorf("Expected to zero initialize integer, but got: %d", *g)
	}
}

func Test_Mutex_constructor(t *testing.T) {
	m := New(func() int { return 1 })

	g, unlock := m.Lock()
	defer unlock()

	if g == nil {
		t.Fatal("Got nil for guarded value")
	}
	if *g != 1 {
		t.Errorf("Expected New to initialize with value 1, but got: %d", *g)
	}
}
