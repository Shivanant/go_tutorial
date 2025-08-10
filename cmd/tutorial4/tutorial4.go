package main

import "fmt"

func main() {

	fmt.Println("This is Tutorial 4!")
	intArr1 := [...]int{2, 4, 5, 6, 67, 4}
	fmt.Println(intArr1)
	fmt.Println(len(intArr1))
	fmt.Println(cap(intArr1))

	intslice1 := []int{2, 4, 5, 6, 67, 4}

	fmt.Println("The slice 1:", intslice1)
	fmt.Println("The slice 1:", len(intslice1))
	fmt.Println(cap(intslice1))

	fmt.Println("Slice cap after inseting a element")
	intslice1 = append(intslice1, 5)
	fmt.Println("The slice 1 cap:", cap(intslice1))
	fmt.Println("The slice 1 len:", len(intslice1))
	fmt.Println(intslice1)

	//ANOTHER WAY OF MAKING A SLICE
	var inslice2 []int = make([]int, 4, 6)
	fmt.Println(inslice2)

	// iterating over a slice

	// for i := range inslice2 {
	// 	fmt.Println(inslice2[i])
	// }

	// for i, val := range intslice1 {
	// 	fmt.Println(i, " val:", val)
	// }

	var map1 map[string]int = map[string]int{"shiva":23,"shiva1": 433}
	for i := range map1{
		fmt.Println(i,map1[i])
	}

	shiva3,ok := map1["shiva3"];
	fmt.Println(shiva3)
	fmt.Println(ok)

	i :=0

	for{
		i+=1
		fmt.Println(i)
		if i ==4 {
			break;
		}
	}

}