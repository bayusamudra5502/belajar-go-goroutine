package app

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T){
	var angka int64
	group := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		group.Add(1)

		go func(){
			defer group.Done()
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&angka, 1)
			}
		}()
	}
}

func TestValue(t *testing.T) {
	var m atomic.Value // Atomic yg bebas mo disimpen apa aja
	group := sync.WaitGroup{}

	m.Store("HAHAHAHAH")

	for i := 0; i < 1000; i++ {
		group.Add(1)
		
		go func(i int) {
			defer group.Done()

			m.Store(fmt.Sprintf("%d", i))
		}(i)
	}

	group.Wait()
	fmt.Println(m.Load())
}