package main 
import "fmt"
func main(){
	var i int = 10 
	var p *int = &i
	// var p *int = new (int)
	fmt.Println(p)
	fmt.Println(*p)

	fmt.Println("Setting the value of *p ")
	*p = 30
	fmt.Println(*p)

	fmt.Println("Tutorial 6 is compleated !")
	fmt.Println(&p)
	fmt.Println(&i)
}