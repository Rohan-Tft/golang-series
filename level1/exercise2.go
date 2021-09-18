package main

import "fmt"

//defining package level variables

var x int
var y string
var z bool


func main(){
	//variables declared but not initialized is always 0
	//in case of int =0 ,string = nill and bool==false|| 0
	//for multiple input use Println
	fmt.Println(x,y,z)
	//output- 0 "" false

	//print f can have single input
	fmt.Printf("x=%T\n",x)
	fmt.Printf("y=%T\n",y)
	fmt.Printf("z=%T\n",z)

}