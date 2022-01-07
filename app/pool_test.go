package app_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPoolSatu(t *testing.T){
	pool := sync.Pool{
		New: func() interface{} {
			return "Objek Baru"
		}, // Ada di structnya. jadi bisa
	}
	wait := sync.WaitGroup{}

	pool.Put("Ayam")
	pool.Put("Bebek")
	pool.Put("Cendol")

	for i := 0; i < 10; i++ {
		wait.Add(1)
		go func(i int){
			// KLo kosong poolnya, nanti kasi nil atau klo 
			// ada New dilakukan fungsi new
			data := pool.Get()
			fmt.Println(i, data)

			time.Sleep(1 * time.Second)
			pool.Put(data)
			wait.Done()
		}(i)
	}

	wait.Wait()
}