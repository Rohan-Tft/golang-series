package main

import "fmt"

//global vs local varibles

//defining global variables
var Fname="Rohan"

func  main()  {
	fmt.Println("Git setup")
	Lname:="Yadav"
	fmt.Println(Fname,"+",Lname)

}