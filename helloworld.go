package main

import "fmt"

// const NAME = "abhay";

var new string = "pom";
//Can be used inside and outside of functions
//Variable declaration and value assignment can be done separately

// newVar := "setting"; 
// we can't use these declaration outside functions i.e. only inside functions
//Variable declaration and value assignment cannot be done separately (must be done in the same line)

func main() {
	// var a,b,c,d,e int = 1,2,3,4,5; - this is possible

	// var student1 string = "abhay";
	// student2 := "om";
	// var student3 = "Jane" //type is inferred

	// fmt.Println(student1);
	// fmt.Println(student2);
	// fmt.Println(student3);
	

	var first string;
	var second int;
	var third bool;
	var fourth float32;

	fmt.Println(first);
	fmt.Println(second);
	fmt.Println(third);
	fmt.Println(fourth);

	var newOne = "adding again";

	const CONST_VAR = "this is a constant and can't be changed"; //Constant names are usually written in uppercase letters for easy identification
	fmt.Print(CONST_VAR," ", newOne);

	//The Printf() Function
	//%v is used to print value - any type
	//%T is used to print the type of Variable
	//%#v - Prints the value in Go-syntax format
	//%% - Prints the % sign
	fmt.Printf("\nthis is a Variable %#v of type %T", newOne,newOne);

}

//go mod init example.com/hello
//go run .\helloworld.go     