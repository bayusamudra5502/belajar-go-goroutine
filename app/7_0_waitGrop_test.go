package app

import (
	"fmt"
	"sync"
	"testing"
)

func TestWitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go meong(i, group)
	}

	group.Wait()
	fmt.Println("Selesaii")
}

func meong(number int, group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	fmt.Printf("Putaran ke-%d\n", number)
}