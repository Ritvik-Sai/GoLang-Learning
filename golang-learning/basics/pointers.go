package main

import "fmt"

func ChangeValue(x *int) {
	*x = 100
}


func ShowPointers( )  {

	var num = 42
fmt.Println("before:", num)
	ChangeValue( &num )

  fmt.Println("after:" , num)
	 var ptr *int = &num
	fmt.Println("ptr points to:", *ptr)

}
