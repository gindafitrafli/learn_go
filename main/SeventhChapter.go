package main

import (
	// "bytes"
	"fmt"
	// "io"
	// "net/textproto"
)

//go always run from main thats why this function is invoked from main
//unlike python with if __name__ == "main"
func learnmethodAndFunctionsAndInterface() {
	// name := "ginda"
	// greet("test", name)
	// fmt.Println(name)//name wont change despite function greet try to modify it
	// a := sum(1,3,4,6,7,8,10)
	// fmt.Println(a)
	// c, error1 := divide(23.34, 543.4) 
	// fmt.Println(c, error1)
	// d, error2 := divide(1.234, 0.0)
	// fmt.Println(d, error2)

	// //anonymous function should be called after being declared
	// for i :=0; i<5; i++ {
	// 	func (i int) {//passing it in case of inconcistency from concurrency 
	// 		fmt.Println(i)
	// 	} (i)//call anonymous function
	// } 
	// //function can also be assigned to variable

	// var div2 func(float64, float64)(float64, error) = func (a, b float64)(float64, error) {
	// 	if b==0.0{
	// 		return 0.0, fmt.Errorf("cannot divide by zero")
	// 	}
	// 	return a/b, nil
	// }
	// d3, err3:= div2(123.345, 456.234)
	// fmt.Println(d3, err3) 

	// //---------------method chapter----------------
	// g := greeter{
	// 	greeting: "test",
	// 	name: "ginda",
	// }
	// g.greet1()
	// g.greet2()
	// fmt.Println(g)


	//interface
	//learnInterface1()
	// learnInterface2()
	// learnInterface3()
	learnInterface4()

}

// type greeter struct {
// 	greeting string
// 	name string
// }

// func (g greeter) greet1() { // this is a method. method is a function that has a value receiver which signifies this method work on defined struct.
// 	fmt.Println(g.greeting, g.name)
// }

// func (g *greeter) greet2() {// passing a pointer will enable us to change property value of greeter
// 	fmt.Println(g.greeting, g.name, "2")
// 	(*g).name = "greet2"
// }

// // if all parameter has the same type we can declare it only at the end(automatically handled by compiler)\
// func greet(msg, name string) {
// 	fmt.Println(msg, name)
// 	name="is"//wont have effect on name variable on parent function, cause in go variable only passing value. if one want to change value on parent function one should pass a pointer to determined variable.
// 	//passing by pointer have huge advantage when passing large object so that large object wont be duplicated. 
// 	//wont have effect on map and slice since those two already pointing to underlying data
// }

// // function in go can have many amount of parameter as long as having same type and located in the end of parameter statement. it will treated as slice in side of that particular function
// func sum(values ... int) (result int){
// 		fmt.Println(values)
// 		for _, v := range values {
// 				result +=v
// 		}
// 		return//returning result as already defined in returned type
// }

// //returning an exception/error in go

// func divide(a, b float64)(float64, error) {
// 	if b==0.0{
// 		return 0.0, fmt.Errorf("cannot divide by zero")
// 	}
// 	return a/b, nil
// }

// func learnInterface1() {
// 	// var w ConsoleWriter = ConsoleWriter{}// we can declare var w like this 
// 	var w Writer = ConsoleWriter{}/* or like this one, but with this one 
// 	we can freely determine precise type of implementor of Writer.*/
// 	/*we declare a type of interface named writer 
// 	and instantiate it with ConsoleWriter as implementor of Writer. 
// 	ConsoleWriter implements Writer through method called wirte, 
// 	this is called implicit implementation. */
// 	w.Write([]byte("hello ginda implicit "))

// 	/*
// 	since go using implicit implementation for a method. implementation can go 2 ways.
// 	the first one is the ordinary one like java where you have several object(struct) 
// 	and you want them to be grouped based on behavior (polymorphism, where one method has different implementation based on what object).
// 	the second one is you have several object from remote package and you want to test with mock, you can create an interface that replicate the method signature
// 	then go with its compiler will automatically determined the implementor.
// 	*/
// }

// type Writer interface {//Single method interface naming convention: method+(er)
// 	Write ([] byte) (int, error)
// }

// type ConsoleWriter struct {}

// func (cw ConsoleWriter) Write(data []byte) (int, error) {
// 	n, err := fmt.Println(string(data))
// 	return n, err
// }

// func learnInterface2() {
// 	 myInt := IntCounter(0)
// 	 var inc Incrementer = &myInt
// 	 for i:=0; i<10; i++ {
// 		fmt.Println(inc.Increment())
// 	 }
// }

// type Incrementer interface {
// 	Increment() int
// }

// type IntCounter int

// func (ic *IntCounter) Increment() int {
// 	*ic++
// 	return int(*ic) /* a bit weird decision by go that you can increment dereferencing pointer 
// 	IntCounter but cannot return dereferencing operator of IntCounter pointer
// 	*/
// }

// func learnInterface3() {
// 	// var wc WriterCloser = NewBufferedWriterCloser()
// 	// wc.Write([]byte("ginda test writer closer"))
// 	// wc.Close()//to close stream buffer

// 	// bwc := wc.(*BufferedWriterCloser)//type casting
// 	// r, ok := wc.(io.Reader)//will error cause dont implement same method
// 	// if ok {
// 	// 	fmt.Println(r)
// 	// } else {
// 	// 	fmt.Println("failed conversion")
// 	// }
// 	// fmt.Println(bwc, *bwc)

// 	var myObj interface{} = NewBufferedWriterCloser()//this is empty interface
// 	if wc2, ok2:=myObj.(WriterCloser); ok2 {
// 		wc2.Write([]byte("ginda test using empty interface"))
// 		wc2.Close()
// 	}

// 	r3, ok3 := myObj.(io.Reader)
// 	if ok3 {
// 		fmt.Println(r3)
// 	} else {
// 		fmt.Println("failed conversion")
// 	}

// 	//type switch
// 	var iType interface{} = 0
// 	switch iType.(type) {
// 	case int: 
// 		fmt.Println("is  integer")
// 	default:
// 		fmt.Println("other type")
// 	}
// }

// type Writer interface {
// 	Write ([]byte) (int , error)
// }

// type Closer interface {
// 	Close () (error)
// }

// type WriterCloser interface {// a struct (an object) is said to implement this interface if he has Write and Close method
// 	Writer
// 	Closer
// }

// type BufferedWriterCloser struct {
// 	buffer *bytes.Buffer
// }

// func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
// 	n, err:= bwc.buffer.Write(data)
// 	if err != nil {
// 		return 0, err
// 	}
// 	v:=make([]byte, 8)
// 	for bwc.buffer.Len() > 8 {
// 		_, err := bwc.buffer.Read(v)
// 		if err != nil {
// 			return 0, err
// 		}
// 		_, err = fmt.Println(string(v))
// 		if err != nil {
// 			return 0, err
// 		}
// 	}
// 	return n, nil
// }

// func (bwc *BufferedWriterCloser) Close() error {
// 	for bwc.buffer.Len() > 0 {
// 		data := bwc.buffer.Next(8)
// 		_, err := fmt.Println(string(data))
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func NewBufferedWriterCloser() *BufferedWriterCloser {
// 	return &BufferedWriterCloser{
// 		buffer: bytes.NewBuffer([]byte{}),
// 	}
// }

func learnInterface4() {
	var wc WriterCloser = &myWriterCloser{}
	/*
	on go the method receiver and the method itself is kind of a map.
	declaring a type with a value only works with value receiver method
	but declaring with a pointer will work both on value receiver method or pointer reciever method.
	*/
	fmt.Println(wc)
}

type Writer interface {
	Write ([]byte) (int , error)
}

type Closer interface {
	Close () (error)
}

type WriterCloser interface {// a struct (an object) is said to implement this interface if he has Write and Close method
	Writer
	Closer
}

type myWriterCloser struct {}

func (mwc myWriterCloser) Write(data []byte) (int, error){
	return 0, nil
}

func (mwc myWriterCloser) Close() error {
	return nil
}
