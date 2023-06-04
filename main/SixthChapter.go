package main

import(
	"fmt"
)

func learnPointer() {
	fmt.Println("sixthChapter")
	//reassigning variable as long as other than map, channel, interface, and function only duplicates its value.
	// a:=123
	// b:=a
	// fmt.Println(a, b)
	// a=2
	// fmt.Println(a, b)

	// // two variables will be connected once it become pointer
	// var c int = -9432
	// var d *int = &c // declaring pointer in complete way must use data type. now d is a pointer
	// fmt.Println(&c, d) //& will refer to the pointer of variable (& is a reference operator)
	// //to display its value you need * operator on variable name
	// fmt.Println(c, *d)
	// c = 22332
	// fmt.Println(c, *d)
	// *d = 293
	// fmt.Println(c, *d)//will change to 293

	// //go wont let you do pointer arithmetic
	
	// e := [3]int{2,3,45}
	// f := &e[0]
	// g := &e[1]
	// fmt.Printf("%v %p %p \n", e, f, g)//distance between element is equal to each element size(in this case int32 is 4 byte)

	// // h := [3]int{2,3,45}
	// // i := &h[0]
	// // j := &h[1]-4
	// // fmt.Printf("%v %p %p \n", h, i, j)//go wont let us to do pointer arithmetic, but still able to do with 'unsafe' package
	
	// var ps pointerStruct = pointerStruct{foo: 42}
	// fmt.Println(ps)

	// //in go we can also declare pointer of a struct
	// var ps2 *pointerStruct = &pointerStruct{foo: 423}
	// fmt.Println(ps2)// will print &{<foo value>} which signifies ps2 is an address which holds pointer with <foo value>

	// var ps3 *pointerStruct = new(pointerStruct)
	// fmt.Println(ps3)// new syntax declaring an address of a struct with default value. New can't be passed an object declaration syntax 

	// var ps4 *pointerStruct
	// fmt.Println(ps4)//an undeclared pointer variable should be nill
	
	// //we can access struct field from a pointer struct with either of following way
	// (*ps3).foo = 123//parentheses needed to prevent go to dereferencing ps4.foo
	// fmt.Println(*ps3)
	// ps3.foo = 1234//parentheses is not needed since it handled by compiler
	// fmt.Println(*ps3)

	// ps4.foo = 12345// can not be done to ps4 since ps4 is nill
	// fmt.Println(*ps4)
	
	




}

type pointerStruct struct{
	foo int
}

