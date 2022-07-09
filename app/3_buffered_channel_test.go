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

	// channel := make(chan string)
	// Channel diatas sama aja kek channel dengan capacity 0

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
		fmt.Println("SATU")
		fmt.Println(<-channel)
		fmt.Println("DUA")
		fmt.Println(<-channel)
		fmt.Println("TIGA")
		fmt.Println(<-channel)
		fmt.Println("EMPAT")
		fmt.Println(<-channel)
		fmt.Println("LIMA")
		fmt.Println(<-channel)
	}()

	fmt.Println("Haha.. ")
	time.Sleep(15 * time.Second)
}
