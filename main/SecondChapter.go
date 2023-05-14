package main

import(
	"fmt"
)

type Doctor struct {
	number int
	actorName string
	companions []string
}

func learnStruct () {
	/**
	====================STRUCT=============================
	*/
	fmt.Println("testStruct")


	drGinda := Doctor{number: 3, actorName: "Ginda", companions: []string{"Lizst", "Bach"}}

	fmt.Println(drGinda)

	//anonymous struct:
	bDoctor:= struct{name string}{name: "Jane Doe"}
	fmt.Println((bDoctor))
	//struct isnt reference type so reassigning to another struct dontaffect the original struct
	cDoctor:= bDoctor
	cDoctor.name = "Jali"
	fmt.Println(cDoctor)
	fmt.Println(bDoctor)
}