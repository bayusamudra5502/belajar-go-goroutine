package app_test

import (
	"fmt"
	"sync"
	"testing"
)

// Map aman dari balapan

func TestMap(t *testing.T){
	data := sync.Map{}
	group := sync.WaitGroup{}

	for i := 0; i < 20; i++ {
		group.Add(1)
		go func(i int) {
			defer group.Done()
			data.Store(i, i*3)
		}(i)
	}

	group.Wait()

	fmt.Println(data.Load(10))
	fmt.Println(data.Load(25)) // nil
	fmt.Println(data.LoadAndDelete(12))

	fmt.Println(data.LoadOrStore(192, 100))
	fmt.Println(data.LoadOrStore(15, 100))

	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true // Klo false dia bakal berenti iterasi
	})
}