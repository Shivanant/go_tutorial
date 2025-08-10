package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
var wg = sync.WaitGroup{}
var items [5] string= [5]string{"apple", "banana","mango", "grape","cherry"}
var items2 []int=[]int{2,3,4,4,6,3}
func main() {
	t0 := time.Now()
	fmt.Println("thisi is items :",items)
	items2 = append(items2, 3)
	fmt.Println("this  is items2 :",items2)

	for i:=0;i<len(items2);i++{
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()

	fmt.Println("The Total Execution time is :", time.Since(t0))
	fmt.Println("This is the end of tutorial")
}

func dbCall(i int){
	delay := rand.Float64()*1000
	time.Sleep(time.Duration(delay)*time.Microsecond)
	fmt.Println("The item from database is :", items2[i])
	wg.Done()
}