package app_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsync(g *sync.WaitGroup, num int){
	defer g.Done()

	
	fmt.Println("Haii,", num)
	time.Sleep(time.Second)
}

func TestWaitGoroutine(t *testing.T) {
	group := &sync.WaitGroup{}
	
	for i := 0; i <= 200 ; i++ {
		group.Add(1)
		go RunAsync(group, i)
	}


	group.Wait()
	fmt.Println("Done!!")
}

func TestOneFunction(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}
	sum := 0

	for  i := 0; i < 100 ; i++ {
		group.Add(1)

		go func() {
			once.Do(func() {
				sum++
			}) // Hanya fungsi tanpa param yang bisa
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println(sum)
}
