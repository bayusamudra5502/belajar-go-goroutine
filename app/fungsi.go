package app

import (
	"fmt"
	"time"
)

// Fungsi target goroutine gaboleh mengembalikan nilai
func Belajar(){
	fmt.Println("Halo, Kita tunggu")
	time.Sleep(1 * time.Second)
	fmt.Println("Selesai")
}

func TampilkanAngka(angka int){
	fmt.Println("Angka", angka)
}

func HitungSaja(angka int) int {
	return angka * 2
}

/* Klo ini hanya bisa ngeluarin data aja (recieve-only) */
func ChannelKeluarSaja(channel <-chan string){
	fmt.Println("Ini cuma bisa ngeluarin data")
	data := <- channel

	fmt.Println(data)
}

// Klo gini artinya channelnya cuma bisa masukin data (Send-only)
func NamaKu(channel chan<- string, input string){
	channel <- "Halo, " + input
}

// Ini channelnya bisa mengirim dan menerima data
func TestLainnya(channel chan string){
	channel <- "Haii, ini pesan dari TestLainnya"
}