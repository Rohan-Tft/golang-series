package main

import "fmt"

//defining not initializing
var sum int64
var isStarted bool
var name string
func main()  {
	 	//go automatically assign unassigned variables that are integer =value 0
		fmt.Println("not initializing sum=",sum)

		//for string or bool 
		//string-"" and bool=true[0]
		fmt.Println("not initializing bool isStarted=",isStarted)
		fmt.Println("not initializing string name=", name)

		//nil for un unitializing printers,functions,interface,slices,channels,maps
}