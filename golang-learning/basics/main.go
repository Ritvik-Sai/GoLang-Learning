package main

import (
	"fmt"
)

func main()  {

	//variable test
fmt.Println("running variable test")
	ShowVars()

	//function test
	fmt.Println("\n running function test")
		result := Add(5 , 3)
	 fmt.Println("Add result is:" , result)
	SayHello("GoLang")

	//pointers test
fmt.Println("\n running pointers test")
ShowPointers()

	//structs test
fmt.Println("\n running structs test")
 ShowStructs()

}
