# go-bettersync

A tiny amount of generic code that improves the builtin `sync.Mutex` type.


## Usage
```
package example

import "github.com/VolatileDream/go-bettersync"

// Your data stracture
type WrappedTypeExample struct {
	content int
}

// In some other type, or function.
func main() {
	var l *bettersync.Mutex[WrappedTypeExample]

	// bettersync.Mutex only exposes the ability to unlock the mutex
	// after it has been locked. This helps to force the use of defer
	// to ensure the mutex is locked and unlocked correctly. A call
	// to Mutex.Lock must always be accompanied with a deferred call
	// to unlock.
	ptr, unlock := l.Lock()
	// Always call unlock!
	defer unlock()

	// do stuff with lock content.
	ptr.content += 1
}
```
