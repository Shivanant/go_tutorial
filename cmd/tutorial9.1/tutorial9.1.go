package main

import (
	"fmt"
	"math/rand"
	"time"
)

const maxVideo float32= 6
const maxVideo2 float32= 3

func main() {
	websites := []string{"google.com", "youtube.com", "duckduckgo.com"}
	goChannel := make(chan string)
	goChannel2 := make(chan string)
	for i:=  range websites{
		go findVideo(websites[i], goChannel)
		go findVideo2(websites[i], goChannel2)
	}
	sendMessage(goChannel,goChannel2)
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
func findVideo2( website string, c chan string){
	for{
		time.Sleep(time.Second)
		var video = rand.Float32()*20
		if video<= maxVideo2{
			c<-website
			break
		}
	}
}

func sendMessage(c1 chan string, c2 chan string){
	select{
	case website := <-c1:
		fmt.Println("the website is from chan1 :", website)
		case website := <-c2:
		fmt.Println("the website is from chan2 :", website)
	
	}
}