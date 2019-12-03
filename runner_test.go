package runner

import (
	"testing"
	"time"
)

func TestNewRunner(t *testing.T) {
	r := New(10, 10*time.Millisecond)
	for i := 0; i < 100; i++ {
		tmp := i
		r.Add(&Task{Do: func() error {
			println(tmp)
			time.Sleep(500 * time.Millisecond)
			return nil
		}})
	}

	time.Sleep(5 * time.Second)
}
