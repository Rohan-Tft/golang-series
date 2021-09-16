package main

import "fmt"

//2 ways of defining variables
var quote="shaken, not stirred"
var sayGreet string ="Hello User"


//global vs local varibles
//name:='rohan'   this is local and has limited scope
//defining global variables
var Fname="Rohan"
//scope of this variable is global and can be useed anywhere

// go is static no dynamic means
//we cannot change type of variables  after intializing them 

func  main()  {
	fmt.Println("Git setup")
	Lname:="Yadav"
	fmt.Println(Fname,"+",Lname)

	//type of vatiables
	fmt.Println("%T",Lname);

	//using `` we van print even space between then as we have given

	//fmt.Println() can have n number of paramerter that we can print
	fmt.Println(quote,"--",sayGreet)
	fmt.Println(`Hello 

		justin, lets start.
		`)

}