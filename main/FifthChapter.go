package main

import(
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func learnControlFlow () {
 /*
	===========================defer===========================
 */
//  fmt.Println("start")
//  fmt.Println("middle")
//  fmt.Println("end")
//  //defer will delay execution until other execution is finished
//  fmt.Println("start")
//  defer fmt.Println("middle") // will resulting in start end middle
//  fmt.Println("end")

//  //defer is lifo based so the first one that defer declared will executed last
//  defer fmt.Println("start")
//  defer fmt.Println("middle") // will resulting in end middle start
//  defer fmt.Println("end")

//  //defer will only delay execution, of the line not delay compile
//  a:="start"
//  defer fmt.Println(a) //will print start, not end
//  a="end"
//  fmt.Println("test")

res, err := http.Get("http://google.com/robots.txt")
if err != nil {
	log.Fatal(err)
}
// robots, err := ioutil.ReadAll(res.Body)
ioutil.ReadAll(res.Body)
res.Body.Close()

if err!=nil {
	log.Fatal(err)
}
// fmt.Printf("%s", robots)
	
 /*
	panic
 */
	//is the same as throw expression in java

	fmt.Println("start")
	// panic("test ginda")
	
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Test Ginda"))
	// })
	// err2 := http.ListenAndServe(":8080", nil)
	// if err2!=nil{
	// 	panic(err.Error())
	// }
 /*
	recover
 */

 // recover acting like catch in other language

 panicker()

}

func panicker() {
	fmt.Println("starts to panic")
	defer func(){
		if err3:=recover(); err3!=nil{//kalo di re panic error nya nanti ada tulisan recovered
			log.Println("Error: ", err3)
			panic(err3)
		}
	}()
	panic("test recover ginda")
}