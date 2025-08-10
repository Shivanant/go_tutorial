package main
import "fmt"

type person struct {
	firstName string
	lastName  string
}

func (e person) fullName() string {
	return e.firstName + " " + e.lastName
}

type gasEngine struct{
	mpg uint 
	gallons  uint 
}

type electricEngine struct {
	mpkwh uint
	kwh uint
}

func (e gasEngine) milesLeft() uint {
	return  e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint {
	return  e.kwh * e.mpkwh
}

type engine interface{
	milesLeft() uint    // must have same function and the 
} 

func canMakeIt(e engine, miles uint){
	if miles <= e.milesLeft() {
		fmt.Println("Yes You can make it !")
	}else{
		fmt.Println("No you cannot make it!")
	}
}

func main() {

	shiva := person{"shivanant", "bhagat"}
	fmt.Println(shiva.firstName)

	shiva2 := person{"shiva", "2"}
	fmt.Println(shiva2.lastName)

	fulnameShiva := shiva.fullName()
	fmt.Println(fulnameShiva)

	var myEngine gasEngine = gasEngine{10, 10}
	fmt.Println(myEngine.milesLeft())

	var electricEngine electricEngine = electricEngine{20, 12}

	canMakeIt(myEngine, 100)
	canMakeIt(electricEngine, 100)
}
