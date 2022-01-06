package app_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCondition(t *testing.T){
	cond := sync.NewCond(&sync.Mutex{})
	group := sync.WaitGroup{}

	defer group.Wait()

	for i := 0; i < 10; i++ {
		group.Add(1)

		go func(i int){
			cond.L.Lock()
			cond.Wait()
			fmt.Println("Done",i)
			cond.L.Unlock()
			group.Done()
		}(i)

	}

	// Kasi sinyal biar bisa jalan
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			cond.Signal()
		}
	}()
}

func TestConditionBroadcast(t *testing.T){
	cond := sync.NewCond(&sync.Mutex{})
	group := sync.WaitGroup{}

	defer group.Wait()

	for i := 0; i < 10; i++ {
		group.Add(1)

		go func(i int){
			cond.L.Lock()
			cond.Wait()
			fmt.Println("Done",i)
			cond.L.Unlock()
			group.Done()
		}(i)

	}

	// Kasi sinyal biar bisa jalan
	go func() {
		time.Sleep(time.Second)
		cond.Broadcast() // Jalanin semua aja
	}()
}