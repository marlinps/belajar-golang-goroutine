package belajar_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	// channel <- "Marlin" //TODO: mengirim data dichannel

	// data := <-channel //TODO: menerima data

	// fmt.Println(data)

	// close(channel) //TODO: tutup channel dengan close

	//TODO: saat membuat channel pastikan ada yg mengirim atau yg menerima data kalau tidak maka akan terjadi error

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Marlin Purnama Sari"
		fmt.Println("Selesai Mengirim Data ke Channel")

	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)

}
