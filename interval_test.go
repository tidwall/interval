package interval

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestInterval(t *testing.T) {
	var count int32
	iv := Set(func(t time.Time) {
		atomic.AddInt32(&count, 1)
	}, time.Millisecond)
	time.Sleep(time.Millisecond * 100)
	iv.Clear()
	iv.Clear()
	if atomic.LoadInt32(&count) == 0 {
		t.Fatal("expected a non-zero value")
	}
	time.Sleep(time.Millisecond)
}
