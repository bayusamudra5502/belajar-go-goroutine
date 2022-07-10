package app_test

import (
	"fmt"
	"runtime"
	"testing"
)

func TestSatuProc(t *testing.T) {
	// Klo tidak negatif, bisa ubah jumlah thread
	// Defaultnya jumlah core
	totalTh := runtime.GOMAXPROCS(-1)
	fmt.Println(totalTh)

	fmt.Println(runtime.NumCPU())

	// Setidaknya ada  go routeine garbace collection
	// sama alur testingnya
	fmt.Println(runtime.NumGoroutine()) 
}