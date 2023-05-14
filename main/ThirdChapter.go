package main

import(
	"fmt"
)

func learnConditional () {
	/**
	====================IF=============================
	*/
	//if in go must have curly braces unlike java when you have only one line

	//if can have initializer syntax like for loop in c++. initializer is ended with semicolon. variable declared in initializer only available in that if block in particular
	// statePopulations := map[string]int{
	// 	"Cali": 123,
	// 	"Texas": 235,
	// }

	// if _, ok := statePopulations["Cali"]; ok {
	// 	fmt.Println("A")
	// }

	// if _, ok := statePopulations["Florida"]; ok {
	// 	fmt.Println("B")
	// }

	/**
	====================SWITCH=============================
	*/

	// a:="a"
	
	// switch a {// can also hase initializer like i:=2+3; i. value of i(the one behind semi colon which already have operation 2+3) is the one will be pased in to case
	// 	case "a", "b": // this is the same with case "a": case "b" in java
	// 		fmt.Println(a)
	// 	case "c", "d": //switch case must be unique and cant be overlaped/same value with other switch branch (if you have a tagged syntax. its okay if you have tagless syntax)
	// 		fmt.Println("cd")
	// 	default:
	// 		fmt.Println("other value for variable a")
	// }
	
	i:=20
	
	switch{
	case i<=20://switch in go has internal break unlike other language. but you still can use it 
			fmt.Println("less than twenty")
			//using command "fallthrough" we can execute next branch as long as the no matter condition is right or false
		fallthrough//but i think this feature kinda like anti pattern in logical control
	case i<=50:
			fmt.Println("less than fifty")
		
	default:
			fmt.Println("other val")
		
	}
	

}