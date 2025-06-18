package main

import "fmt"

type person struct {
	name string
	age int
}


func ShowStructs( ) {

	p1 := person{ "X" , 22 }
	  fmt.Println("person 1:", p1)
	fmt.Println( "name is", p1.name )
fmt.Println("age is",p1.age)

	// different type
	var p2 person
	 p2.name = "Y"
p2.age = 30

	fmt.Println("p2 =>", p2)

	// another way
	 p3 := person{}
		p3.name = "Z"
	p3.age = -1

	fmt.Println("3rd person",p3)
}
