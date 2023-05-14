package main

import (
	"fmt" //go also support aliasing import
	// "rsc.io/quote"
)
//naming convention in go
//var k float32 = 26 //==>visible to package main only
//var K  //==> visible globally
//accronym  theURL is correct, theUrl is incorrect 
func main() {
	// learnStruct()
	// learnConditional()
	// learnLoop()
	learnControlFlow()
	fmt.Println("main")

	// var i = 43
	// fmt.Println(i+1)
	// j := 69
	// fmt.Printf("%v, %T \n", j, j)
	// fmt.Println("test Ginda")
	// fmt.Println(quote.Go())

	// var k float32 = 26//local variable
	// fmt.Printf("%v, %T \n", k, k)

	// var l float64 = 244
	// fmt.Printf("%v, %T \n", l, l)

	// var o = 24.
	// fmt.Printf("%v, %T \n", o, o)

	// grades := [3]int{1,2, 3}

	// fmt.Printf("Grades: %v\n", grades)

	// gradesb := [...]int{12,12,34}

	// fmt.Printf("Gradesb: %v\n", gradesb)

	// var students [3]string
	// fmt.Printf("Students: %v\n", students)

	// students[0]="test"
	// students[1]="23"
	// students[2]="221"

	// fmt.Printf("Students: %v\n", students)

	// students[2]="123"

	// fmt.Printf("Students: %v\n", students)

	// //if you assign an array to another array you only copy its value. like c to have reference between each array we use pointer

	// a := [...]int{333,2,100}
	// b := a
	// b[1]= 5
	// fmt.Println(a)
	// fmt.Println(b)

	// c := [...]int{1,2,4}
	// d := &c
	// d[1]= 5
	// fmt.Println(c)
	// fmt.Println(d)

	//SLICE
	// e :=[] int {2,434,577}
	// f := e[:2]//if f = e[:] f still referring to e so. that if e change element value f value will also the same as e unlike ptthon
	// f[1] = 444
	// fmt.Println(e)
	// fmt.Println(f)
	// fmt.Println(len(e))
	// fmt.Println(cap(e))

	// g:=make([]int, 3, 100)
	// fmt.Println(len(g))
	// fmt.Println(cap(g))
	// fmt.Println(g)
	// g = append(g, 100)
	// fmt.Println(g)

	// g = append(g, []int{2313, 42,464}...)//slice can be spread, unlike array
	// fmt.Println(g)	
	// h :    = append(f[:], g[:]...)
	// fmt.Println(h) 
	
	/*
	==============MAP=====================
	*/

	// statePopulations := map[string]int{
	// 	"California": 23123123,
	// 	"Texas": 4309,		
	// }
	// fmt.Println(statePopulations)
	// fmt.Println(statePopulations["Texas"])
	// fmt.Println(statePopulations["Ginda"])
	// statePopulations["Ginda"]=213123
	// fmt.Println(statePopulations)
	// delete(statePopulations,"Texas")
	// fmt.Println(statePopulations)
	// delete(statePopulations,"Texas")
	// fmt.Println(statePopulations)
	// _, ot :=statePopulations["Ohio"] //second parameter return parameter of getting value of key something behave like isExist() in java
	// fmt.Println(ot)

	/*
		MAP is also reference type so if 2 statePopultions map is assigned to other variable.
		 and then modify the original map the other var will also changed
	*/

	
	
}