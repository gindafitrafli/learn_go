package main

import(
	"fmt"
	"time"
	"sync"
	// "runtime"
)

var wg = sync.WaitGroup{}
// var counter = 0
var m = sync.RWMutex{}

func learnConcurrency() {
	fmt.Println("")
	// testSimpleGoRoutine()
	time.Sleep(100*time.Millisecond)
	// go testSimpleGoRoutine()//this line will executed as long as main thread execution wasnt finished so time.sleep is necessary
	/*using 'go' keyword allows you to use a 'green' thread 
	(for particular execution) that is differ than other  execution*/
	/*
	most programming language use OS thread, so they use callstack that will be passed into that thread.
	it tends to be very large in memory and also takes long time to start and destroy a thread. thats why ther is a thread pool.

	this is differ with go which abstract a thread using go routine. 
	inside go runtime there is a scheduler, that will mapped each go routine on to OS thread for periods of time.
	And scheduler will take turn for every available thread and assign every go routine certain amount run time but we dont have to interact with those low level threads directly.
	with this abstraction go routine can be very small in size and reallocated very quick to create and destroy.
	*/

	//==========================details on how the go routine execution works=======================
	// executionOrderMessage := "first declaration"
	// go func(){
	// 	fmt.Println(executionOrderMessage)
	// } ()//will print out second declaration.this condition is called race condition should be avoided. so we should pass a parameter to this anon function
	// go func(msg string){
	// 	fmt.Println(msg)//will print out first declaration
	// }(executionOrderMessage)
	// executionOrderMessage = "second declaration"
	// time.Sleep(100*time.Millisecond)// this is not appropriate, can induce race condition we can use wait group
	
	//==========================details on how the go routine execution works with proper code using waitgroup=======================
	//waitgroup synchronize all routines together
	// msg:="Hello"
	// wg.Add(1)//adding one go routine to waitgroup. starting value is 0
	// go func (msg string)  {
	// 	fmt.Println(msg)
	// 	wg.Done() //to decrement waitgroup value that has been added with add()
	// }(msg)
	// msg="Goodbye"
	// fmt.Println(msg)
	// wg.Wait()//to wait for a routine to execute
	// //but waitgroup cant synchronize 2 routines together that using same data

	//==========================synchronizing 2 routines that using same data=====================
	// for i:=0; i<10; i++ {
	// 	wg.Add(2)
	// 	go sayHello()//
	// 	go increment()//these two routines will race each other so the increment is not fully operate after say hello
	// } 
	// wg.Wait()

	//SO WE USE MUTEX FOR THIS CASE
	for i:=0; i<1; i++ {
	// for i:=0; i<10; i++ {
		// wg.Add(2)
		m.RLock()//needs to be stated outside routines scope
		go sayHello()//
		// m.RUnlock()//needs to be stated inside routines scope
		m.Lock()//needs to be stated outside routines scope
		go increment()//these two routines will race each other so the increment is not fully operate after say hello
		// m.Unlock()//needs to be stated inside routines scope
	} 
	// wg.Wait()

	//but this method is the same as not using go routines at all, because we forcing synchronizing execution
	//even not using go routines is faster because no lock and unlocking process

	//or we can also use runtime.GOMAXPROCS()
	//runtime.GOMAXPROCS(-1)//without any declaration it will defaulted to number of cpu core. -1 is also automatically not a declaration
	//runtime.GOMAXPROCS(1)//declaration of amount of thread processing to 1. will autonamtically run application in single threaded way.
	//runtime.GOMAXPROCS(-1)//returning 1 
	//best practice in runtime.GOMAXPROCS() is using value grater than one to find any race condition, in development stage.
	// and finding optimal number during performance test as adding too many thread will lead to excessive memory usage and burden go runtime to assign each routine to a thread.


}

// func testSimpleGoRoutine() {
// 	fmt.Println("test simple go routine ginda")
// }

func sayHello() {
	
	// fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	
	// wg.Done()//mutex can be run without waitgroup
	
}

func increment() {
	
  // counter++
	m.Unlock()
	// wg.Done()
	
}

func learnChannel() {
	//channel are designed to transfer data between go routines
	
	//to make channel only possible through make statement
	// ch := make(chan int)//channel is strongly type so only declared data type can be passed
	// wg.Add(2)
	// go func() {
	// 	i:=<-ch//this is output of a channel. reveiving data from a channel and assigning it to a variable 'i'
	// 	fmt.Println(i)
	// 	wg.Done()
	// }()
	// go func() {
	// 	i:=54
	// 	ch<-i//this is input of a channel. passing copy of a value to a channel
	// 	i=234//because go always copy a value when passing a variable to an object. changing it afterwards wont impact to the channel.
	// 	fmt.Println(i)
	// 	wg.Done()
	// }()
	// wg.Wait()

	learnChannel2()
} 

func learnChannel2() {
	//go routine only able to receive value as much it has buffer capacity or as much outflow it has.
	// ch:= make(chan int)
	ch := make(chan int, 100)//declaring in with size over head which make this channel a buffer channel

	// // go func() {
	// // 	i := <- ch
	// // 	fmt.Println(i)
	// // 	wg.Done()
	// // }()//will error unless declaring in with size over head which make this channel a buffer
	// for j:=0; j< 5; j++ {
	// 	wg.Add(2)
	// 	go func() {
	// 		i := <- ch
	// 		fmt.Println(i)
	// 		wg.Done()
	// 	}()//moving this line to outside the loop will causing error
	// 	go func ()  {
	// 		ch<-32//when outflow channel is moved to outside this line will deadlocking the execution
	// 		wg.Done()
	// 	}()
	// }
	// wg.Wait()

	//======channel will automatically look for input and output===============================

	/*
	below execution will look like: 
	after declaration wg.add
	after declaration method atas
	after declaration method bawah
	method atas
	method bawah
	output channel 1: 54
	selesai method atas
	output channel 1: 123123123
	selesai method bawah
	after declaration wg.wait
	*/
	// wg.Add(2)
	// fmt.Println("after declaration wg.add")
	// // go func(ch<- chan int) {//to make routines more clear we declare a routine receiving an output from channel. but there will be error when assigning value to a channel inside of this method
	// go func() {
	// 	fmt.Println("method atas")
	// 	i:=<-ch
	// 	fmt.Printf("output channel 1: %v\n", i)
	// 	ch<-123123123
	// 	fmt.Println("selesai method atas")
	// 	wg.Done()
	// }()//(ch)//need to pass the channel to this method
	// fmt.Println("after declaration method atas")
	// // go func(ch chan<- int) {//to make routines more clear we declare a routine assigning a value to channel. but there will be error when receive value from a channel inside of this method
	// go func() {
	// 	fmt.Println("method bawah")
	// 	i:=54
	// 	ch<-i
	// 	fmt.Printf("output channel 1: %v\n",<-ch)
	// 	fmt.Println("selesai method bawah")
	// 	wg.Done()
	// }()//(ch)//need to pass the channel to this method
	// fmt.Println("after declaration method bawah")
	// wg.Wait()
	// fmt.Println("after declaration wg.wait")


	//=============channel needs to be closed before it can be ranged, otherwise it will error once there is no data left=====================
	wg.Add(2)
	go func (ch<-chan int)  {
		for i:=range ch {
			fmt.Println(i)
		}
		wg.Done()
	} (ch)

	go func (ch chan<- int)  {
		ch<-42
		ch<-234
		close(ch)
		//ch<-123//will error on closed channel
		wg.Done()
	}(ch)
	wg.Wait()
	//i, ok:= <-ch //the 'ok' signify whether a channel is open or closed
	
	//=====================select statement=============
	
	//select statement is to manage go routine that dont have an obvious way to close

}