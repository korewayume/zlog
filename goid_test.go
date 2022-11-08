package zlog

import (
	"sync"
	"testing"
	"time"
)

func dispatchTime(f func()) time.Duration {
	var wg sync.WaitGroup
	s := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			f()
		}()
	}
	wg.Wait()
	return time.Since(s)
}

func testEmpty() {
	for i := 0; i < 1000; i++ {
	}
}

func testGoID() {
	for i := 0; i < 1000; i++ {
		CurGoroutineID()
	}
}

func TestEncrypt(t *testing.T) {
	t.Logf("testEmpty: %s", dispatchTime(testEmpty))
	t.Logf("testGoID: %s", dispatchTime(testGoID))
}
