package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func leaf1(){
	time.Sleep(15*time.Second)
	fmt.Println("1")
}

func leaf2(){
	time.Sleep(3*time.Second)
	fmt.Println("2")
}

func branch(){
	go leaf1()

	go leaf2()

	time.Sleep(5*time.Second)
	fmt.Println("branch")
}
func main() {
	go branch()

	ch:=make(chan os.Signal)
	signal.Notify(ch)
	<-ch
}