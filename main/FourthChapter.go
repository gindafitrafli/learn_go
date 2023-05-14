package main

import(
	"fmt"
)

func learnLoop () {
	fmt.Println("loop")
	// loop in go only available with for. while is not supported by go.
	//for loop  go can only have one iterator unlike c++, but we can achieve that with double assignment

	// for i, j := 0, 0;i<5 && j<5; i, j=i+1, j+1{
	// 	fmt.Println(i+j)
	// }

	//break can be used to breaking group of loop. example of declaring loop group is on line 16
	// Aasd:
	// for i:=0; i<10; i++ {
	// 	for j :=0; j<10; j++ {
	// 		fmt.Println(i,j)
	// 		if j*i>6{
	// 			break Aasd
	// 		}
	// 	}
	// }

	//go can also loop through collection like string map or list using range. 
	//range will return 2 value the iterator and the value itself.
	//it means that for map the iterator will be the key and value is the map value. 
	//for other collection the iterator will be index and the value will be each element in collection

	// loopMap:=map[string]string{
	// 	"bacon": "check",
	// 	"lettuce": "check",
	// 	"tomato": "check",
	// }

	// for k, v := range(loopMap){
	// 		fmt.Println(k, v)
	// }

	// loopString:="Loop String"
	// for k, v := range loopString {
	// 		fmt.Println(k, string(v))
	// }
}