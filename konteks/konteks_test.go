package konteks_test

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func counterKu(ctx context.Context) <- chan int {
	channel := make(chan int)
	
	go func(){
		defer close(channel)
		counter := 0

		for {
			select{
				case <- ctx.Done():
					return
				default:
					channel <- counter
					counter++
					time.Sleep(time.Second) // Simulasi slow
			}
		}
	}()

	return channel
}

func TestKonteksPertama(t *testing.T){
	ctx := context.Background()
	ctx2 := context.TODO() // Ini juga ada ya

	fmt.Println(ctx)
	fmt.Println(ctx2)
}

func TestBuatContext(t *testing.T) {
	ctx := context.Background()

	childA := context.WithValue(ctx, "ID", 1)
	childAA := context.WithValue(childA, "Score", 100)

	childB := context.WithValue(ctx, "ID", 2)
	childBA := context.WithValue(childB, "Score", 50)


	fmt.Println(ctx)
	fmt.Println(childA)
	fmt.Println(childAA)
	fmt.Println(childBA)

	fmt.Println(childA.Value("ID"))
	fmt.Println(childA.Value("Score")) // Gaada walau sampai parent tertinggi

	fmt.Println(childAA.Value("ID")) // Dari parent coy
	fmt.Println(childAA.Value("Score"))
}

func TestCancelContext(t *testing.T) {
	fmt.Println("Total Goroutines", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)
	
	dataStream := counterKu(ctx)

	fmt.Println("Total Goroutines", runtime.NumGoroutine())

	for n := range dataStream {
		fmt.Println("Counter", n)

		if n == 10 {
			break
		}
	}

	cancel()
	<- dataStream // Cleaning aja

	time.Sleep(time.Second/2)

	fmt.Println("Total Goroutines", runtime.NumGoroutine())
}

func TestCancelTimeout(t *testing.T) {
	fmt.Println("Total Goroutines", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5 * time.Second)
	
	defer func() {
		cancel()

		time.Sleep(time.Second/2)
		fmt.Println("Total Goroutines", runtime.NumGoroutine())
		}()
		
		dataStream := counterKu(ctx)
		
	fmt.Println("Total Goroutines", runtime.NumGoroutine())
	
	for n := range dataStream {
		fmt.Println("Counter", n)
		
		if n == 10 {
			break
		}
	}
	
	<- dataStream // Cleaning aja
	fmt.Println("Total Goroutines", runtime.NumGoroutine())
}

func TestDelayTimeout(t *testing.T) {
	fmt.Println("Total Goroutines", runtime.NumGoroutine())
	parent := context.Background()

	// WithDeadline tu ditentuin kapan deadlinenya
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(time.Second * 5))
	
	defer func() {
		cancel()

		time.Sleep(time.Second/2)
		fmt.Println("Total Goroutines", runtime.NumGoroutine())
		}()
		
		dataStream := counterKu(ctx)
		
	fmt.Println("Total Goroutines", runtime.NumGoroutine())
	
	for n := range dataStream {
		fmt.Println("Counter", n)
		
		if n == 10 {
			break
		}
	}
	
	<- dataStream // Cleaning aja
	fmt.Println("Total Goroutines", runtime.NumGoroutine())
}