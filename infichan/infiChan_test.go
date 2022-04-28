package infiChan

import (
	"sync"
	"time"

	"testing"
)

func TestInfiChan1(t *testing.T) {
	ic := New(1)
	//wg sync.WaitGroup
	go func() {
		for i := 0; i < 1000; i++ {
			ic.I <- i
		}
		//time.Sleep(1 * time.Second)
		ic.Close()
	}()

	var num = 0
	for r := range ic.O {
		if r != num {
			t.Error("failed ,want to get ", num, " but ", r)
		}
		num++
	}
	if num != 1000 {
		t.Error("the date lose ", num)
	}
}

func TestInfiChan2(t *testing.T) {
	ic := New(4)
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(num int) {
			for i := num; i < num+125; i++ {
				ic.I <- i
			}
			wg.Done()
		}(i * 125)
	}

	wg.Wait()
	//time.Sleep(time.Millisecond)

	if ic.Len() != 375 {
		t.Error("failed ,want to get the num of data : 375 , but ", ic.Len())
	}
	var num = 0
	for num < 125 {
		<-ic.O
		num++
	}
	time.Sleep(time.Millisecond)
	if ic.Len() != 250 {
		t.Error("failed ,want to get the num of data : 250, but ", ic.Len())
	}
	ic.Close()
}
