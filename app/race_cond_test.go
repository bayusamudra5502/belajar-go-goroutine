package app_test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T){
	channel := make(chan int)
	defer close(channel)

	go func(){
		x := 0

		for i := 0; i < 1000; i++ {
			go func(){
				for i := 0; i < 100; i++ {
					x++
				}
			}()
		}

		time.Sleep(1 * time.Second)
		channel <- x
	}()

	fmt.Println("Hasilnya adalah", <-channel) // Bisa ga 100.000, balapan soalnya 
}

func TestRaceConditionMutex(t *testing.T) {
	// Solusi dari masalah sebelumnya, pake mutex!
	channel := make(chan int)
	defer close(channel)
	
	go func(){
		var mutex sync.Mutex
		x := 0

		for i := 0; i < 1000; i++ {
			go func(){
				for i := 0; i < 100; i++ {
					mutex.Lock()
					x++
					mutex.Unlock()
				}
			}()
		}

		time.Sleep(1 * time.Second)
		channel <- x
	}()

	fmt.Println("Hasilnya adalah", <-channel) // Pasti 100.000
}

type NumberKu struct {
	RWMutex sync.RWMutex
	Value int
}

func (b *NumberKu) Add(x int){
	b.RWMutex.Lock()
	b.Value += x
	b.RWMutex.Unlock()
}

func (b *NumberKu) Get() int {
	b.RWMutex.RLock()
	hasil := b.Value
	b.RWMutex.RUnlock()
	return hasil
}

func TestMutexRW(t *testing.T) {
	x := NumberKu{}

	for i := 0; i < 100; i++ {
		go func(){
			for i := 0; i < 100; i++ {
				x.Add(1)
				fmt.Println(x.Get())
			}
		}()
	}

	
	time.Sleep(time.Second)
	fmt.Println("Hasilnya", x.Value)
}


func TestAtomic(t *testing.T) {
	group := sync.WaitGroup{}
	var cnt int64 = 0

	for i := 0; i < 100; i++ {
		group.Add(1)

		go func(){
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&cnt, 1) // Sehat akan balapan
				
				// value := atomic.LoadInt64(&cnt)
				// value += 2
				// atomic.StoreInt64(&cnt, value)
			}

			group.Done()
		}()
	}

	group.Wait()
	fmt.Println(cnt)
}