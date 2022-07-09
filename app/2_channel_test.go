package app_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/bayusamudra5502/belajar-go-goroutine/app"
)

// Channel berjalan sinkron. Jika ada data di channel,
// Alur eksekusi akan terhenti atau jika data blom sampai
// Datanya, dia juga nunggu tuh

func TestChannelAwal(t *testing.T){
	channel := make(chan string)
	defer close(channel) // JANGAN LUPA!

	// Saat mainin sleepnya, lihat timenya berapa lama

	go func() {
		time.Sleep(3  * time.Second) // Coba pindahin ini sebelum channel ama sesudah
		channel <- "Ini pesan channelku"
		fmt.Println("Setelah Channel aku")
	}()
		
	fmt.Println(<- channel)
}

func TestChannelMuter(t *testing.T){
	channel := make(chan string)
	defer close(channel)

	go app.NamaKu(channel, "Bayu")
	data := <- channel

	fmt.Println(data)
}

func TestChannelV2(t *testing.T){
	channel := make(chan string)
	defer close(channel)

	go app.TestLainnya(channel)
	data := <- channel

	fmt.Println(data)
}

func TestLainnya(t *testing.T){
	channel := make(chan string)
	defer close(channel)

	fmt.Println(len(channel))
	fmt.Println(cap(channel))

	go app.ChannelKeluarSaja(channel)
	channel <- "Ini Kucing"
	go app.ChannelKeluarSaja(channel)
	channel <- "Meok" // Harus dikonsumsi dulu baru bisa jalan lagi

	fmt.Println("Done")
}