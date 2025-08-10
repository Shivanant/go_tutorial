package main

import (
	"fmt"
	"math/rand"
	"time"
)

const maxVideo float32= 6

func main() {
	websites := []string{"google.com", "youtube.com", "duckduckgo.com"}
	goChannel := make(chan string)
	for i:=  range websites{
		go findVideo(websites[i], goChannel)
	}
	fmt.Println(<-goChannel)
	fmt.Println("This is the end of tutorial 9.1 !")
}

func findVideo( website string, c chan string){
	for{
		time.Sleep(time.Second)
		var video = rand.Float32()*20
		if video<= maxVideo{
			c<-website
			break
		}
	}
}