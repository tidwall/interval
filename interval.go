package interval

import "time"

// Interval represents a repeatedly called a function with fixed time delay.
// Use Start() to start running an interval.
type Interval struct {
	clear chan bool
}

// Set repeatedly calls a function with a fixed time delay between each call.
// Returns an interval object which can be later removed by calling Clear().
func Set(fn func(t time.Time), delay time.Duration) Interval {
	iv := Interval{clear: make(chan bool)}
	ticker := time.NewTicker(delay)
	go func() {
		for {
			select {
			case <-iv.clear:
				iv.clear <- true
				return
			case t := <-ticker.C:
				fn(t)
			}
		}
	}()
	return iv
}

// Clear cancels the repeating action which was previously established by Set().
// This operation waits for an active action to complete before returning.
// This function should never be called from inside of the repeated function,
// otherwise the app will block.
func (iv Interval) Clear() {
	defer func() {
		recover() // ignore "closed channel" panics
	}()
	iv.clear <- true
	<-iv.clear
	close(iv.clear)
}
