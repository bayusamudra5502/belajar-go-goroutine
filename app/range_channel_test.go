package app_test

import (
	"fmt"
	"testing"
	"time"
)

// Range channel adalah metode untuk mengambil data channel
// terus menerus hingga channelnya ditutup oleh penyedianya

// Cocok klo kita gatau ada berapa data yg harus ditunggu

func TestRange1(t *testing.T){
	channel := make(chan int)

	go func()  {
		for i := 0 ; i < 10; i++ {
			channel <- i
		}

		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)	
	}
}	

func TestRange2(t *testing.T){
	// klo 0 sizenya -> channel biasa
	chGambar := make(chan int, 2) 

	go func()  {
		for i := 0 ; i < 10; i++ {
			chGambar <- i
			fmt.Println("Kirim", i)
		}
		
		close(chGambar)
		}()
		
		for data := range chGambar {
			fmt.Println(data)
			time.Sleep(500 * time.Millisecond)
	}
}