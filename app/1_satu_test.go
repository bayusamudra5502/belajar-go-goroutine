package app_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/bayusamudra5502/belajar-go-goroutine/app"
)

func TestSatu(t *testing.T){
	fmt.Println("Test satu kita mulai")
	go app.Belajar()
	go app.Belajar()
	fmt.Println("Test satu udah jalanin goroutine")

	// Klo ga ditunggu, goroutinenya ga dijalanin
	time.Sleep(2 * time.Second)
}

func TestDua(t *testing.T){
	for i := 0; i < 1000; i++ {
		go app.TampilkanAngka(i)
	}

	time.Sleep(3 * time.Second)
}

func BenchmarkTesting(b *testing.B){
	for i := 0; i < b.N; i++ {
		go app.HitungSaja(i)
	}
}

func BenchmarkTanpaGo(b *testing.B){
	for i := 0; i < b.N; i++ {
		app.HitungSaja(i)
	}
}