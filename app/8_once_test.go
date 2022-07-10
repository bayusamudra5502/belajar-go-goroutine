package app

import (
	"fmt"
	"sync"
	"testing"
)

func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)

		go func(){
			defer group.Done()
			
			once.Do(maung)
		}()
	}

	group.Wait()
	fmt.Println("Berees")
}

func maung() {
	fmt.Println("AING MAUNGG!!!")
}