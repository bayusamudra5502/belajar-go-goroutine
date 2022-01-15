package app_test

import (
	"fmt"
	"testing"
	"time"
)

// Buffered Channel

func TestBuffered1(t *testing.T) {
	channel := make(chan string, 3) // Ga lagi blocking selagi masi ada yg kosong
	// tapi sekalinya udah penuh si channelnya, dia nunggu
	// sampe ada yg kosong baru masuk

	fmt.Println(cap(channel))
	fmt.Println(len(channel))

	go func() {
		channel <- "Kumbang"
		channel <- "Kucing"
		channel <- "Kuda nil"
		fmt.Println("HUPLA 1")
		channel <- "Kuda nil Nunggu"
		fmt.Println("HUPLA 2")

		time.Sleep(8 * time.Second)
		channel <- "Nuggu dulu"
	}()

	go func() {
		time.Sleep(4 * time.Second)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	fmt.Println("Haha.. ")
	time.Sleep(15 * time.Second)
}
