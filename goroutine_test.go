package belajar_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld() // FIXME:menambahkan perintah go (dirun melalui goroutine)
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)

}

func DisplayNumber(number int) {
	fmt.Println("Display", number)

}

//TODO:Testing Goroutine sangatlah ringan, tidak ada memory overflow
func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(10 * time.Second)

}
