package app_test

import (
	"fmt"
	"testing"
	"time"
)

// Select channel kepake klo banyak channel yg ingin dibaca

func panicHandler(){
	recover()
	fmt.Println("Ok")
}

func TestSelectChannel(t *testing.T){
	defer panicHandler()

	channel1 := make(chan int)
	channel2 := make(chan int)

	go func ()  {
		for i := 0; i < 10; i++ {
			channel1 <- i
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for i := 10; i < 20; i++ {
			channel2 <- i
			time.Sleep(time.Second)
		}
	}()

	cnt := 0

	for {
		select {
			case data := <- channel1:
				fmt.Println("Data dari channel 1", data)
			case data := <- channel2:
				fmt.Println("Data dari channel 2", data)
			default: // Klo gaada kesini, daripada nunggu
				fmt.Println("Tunggu dulu.")
				time.Sleep(500 * time.Millisecond)
		}

		if (cnt == 10){
			break
		}

		cnt++
	}

	close(channel1)
	close(channel2)
}
