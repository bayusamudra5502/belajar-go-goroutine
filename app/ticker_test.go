package app_test

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker1(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	isStopped := false

	go func() {
		time.AfterFunc(10 * time.Second,func(){
			ticker.Stop()
			isStopped = true
		})
	}()

	for t := range ticker.C {
		fmt.Println(t)

		if isStopped {
			return // Disini ada memory leak ya!!
		}
	}
}

func TestTicker2(t *testing.T) {
	channel := time.Tick(time.Second)
	isStopped := false

	go func(){
		time.Sleep(5 * time.Second)
		isStopped = true
	}()

	for t := range channel {
		fmt.Println(t)

		if isStopped {
			return
		}
	}	
}
