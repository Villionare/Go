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
	var third bool = true;
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
	//there are soo many types of formatting verbs for everydata type, LEARN WHILE DOING. 
	fmt.Printf("\nthis is a Variable %#v of type %T %v", newOne,newOne, third);

	//Go has three basic data types:
	// bool: represents a boolean value and is either true or false
	// Numeric: represents integer types, floating point values, and complex types
	// string: represents a string value

	//Numeric datatypes 
	//1. Integers - signed (+ve & -ve) || unsigned (+ve)

	//unsigned integers are declared with "uint" keyword.

	var unsignedInterger uint64 = 3445654675567564323;
	fmt.Printf("\n%v is the value and the type: %T", unsignedInterger, unsignedInterger);

	//float are also of 2 types - float32 && float64, nothing is special.
	var floatOne float64 = 34534.456; //this is used to store both decimal and floating point values
	fmt.Println("\n",floatOne);

	//string is just string in Go, nothing else

	//ARRAYS IN Go - two types to decalre:
	
	//1. - with var keyword - var array_name = [length]datatype{values} -- length is predefined in this case
	//OR 
	//var array_name = [...]datatype{values} //here length is inferred - length is defined through the no. of values
	var firstArray = [20]int16{1,2,3,4,5,6,7,8,9};
	var secondArray = [...]int{234,4,6,45,7};
	fmt.Println(firstArray);
	fmt.Println(firstArray[2]);
	fmt.Println(secondArray);
	fmt.Println(secondArray[4]);

	//2. - With the := sign:
	//array_name := [length]datatype{values} // here length is defined
	//OR
	//array_name := [...]datatype{values} // here length is inferred
	thirdArray := []string{"this","is","a","new","array"};
	fmt.Println(thirdArray);

	//it is possible to initialise only specific array indexes.
	arr1 := [5]int{1:10,2:40};
	fmt.Println(arr1);
	fmt.Println("The size of the array is: ",len(arr1));

	
}

//go mod init example.com/hello
//go run .\helloworld.go     