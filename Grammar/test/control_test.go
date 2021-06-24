package test

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

func TestControl(t *testing.T) {
	var ops uint64 = 0
	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)
				fmt.Print()
			}
		}()
	}
	time.Sleep(time.Second * 30)
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
	opsFinal2 := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal2)
}
