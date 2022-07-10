package app_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer1(t *testing.T) {
	timer := time.NewTimer(5 * time.Second) // Dikasi data timer
	last := time.Now()
	
	time := <- timer.C // Time saat channel keluar
	fmt.Println(last)
	fmt.Println(time)
	fmt.Println("Delta :", time.Sub(last))
}

func TestTimer2(t *testing.T) {
	timer := time.After(5 * time.Second) // Dapet data time
	last := time.Now()
	
	time := <- timer
	fmt.Println(last)
	fmt.Println(time)
	fmt.Println("Delta :", time.Sub(last))
}

func TestTimer3(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(time.Second, func() {
		defer group.Done()
		
		fmt.Println("Ini setelah 1 detik")
		fmt.Println("Mirip kayak setTimeout di js")
	})

	group.Wait()
}