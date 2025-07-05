package main

import (
	"fmt"
)

func firstFunction(param string){
	fmt.Println("this is my First Go function" , param);
}

// const NAME = "abhay";

// var new string = "pom";
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

	//SLICES IN Go
	//slices are just like arrays and are more flexible
	//However, unlike arrays, the length of a slice can grow and shrink as you see fit.

	//In Go, there are several ways to create a slice:

	// Using the []datatype{values} format
	// Create a slice from an array
	// Using the make() function
	
	fmt.Println("Creating the first Slice:");
	firstSlice := []int{3,4,5,6,7,7};
	fmt.Println(cap(firstSlice)); //- printing the capacity
	fmt.Println(len(firstSlice)); //- length 
	fmt.Println(firstSlice);
	fmt.Println(firstSlice[2]);
	fmt.Println(firstSlice[1]);
	fmt.Println(firstSlice[0]);

	//Creating a slice from an array:
	//var myarray = [length]datatype{values} // An array
	//myslice := myarray[start:end] - A slice made from the array

	//creating a slice using make() function:- 
	// slice_name := make([]type, length, capacity)
	makeSlice := make([]int, 5, 10);
	fmt.Println(makeSlice);
	fmt.Println(cap(makeSlice));
	fmt.Println(len(makeSlice));

	//Appending elements to a slice:
	//slice_name = append(slice_name, element1, element2, ...)

	//Append one slice to another
	//slice3 = append(slice1, slice2...)
	//Note: The '...' after slice2 is necessary when appending the elements of one slice to another.
	myslice1 := []int{1,2,3}
	myslice2 := []int{4,5,6}
	myslice3 := append(myslice1, myslice2...)

	fmt.Printf("myslice3=%v\n", myslice3)
	fmt.Printf("length=%d\n", len(myslice3))
	fmt.Printf("capacity=%d\n", cap(myslice3))


	// Memory Efficiency
	//When using slices, Go loads all the underlying elements into the memory.
	//If the array is large and you need only a few elements, it is better to copy those elements using the copy() function.
	//The copy() function creates a new underlying array with only the required elements for the slice. This will reduce the memory used for the program. 
	//copy(dest, src)

	fmt.Println("copy() function");
	tempArray := [10]int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(tempArray);

	//creating a slice
	tempSlice := []int{11,22,33,44,55}
	fmt.Println(tempSlice);

	//convert an array into a slice 
	convertedSlice := tempArray[1:5];
	// fmt.Println(convertedSlice);

	//copy the elements of converted slice into temp Slice
	newSlice := make([]int, len(convertedSlice));
	
	copy(newSlice, convertedSlice); //we have put the elements of convertedSlice into newSlice, which eas empty
	fmt.Println(newSlice);

	//Go Operators
	//Arthematic Operators
	//Assignment Operators
	//Comparision Operators
	//Logical Operators
	//Bitwise Operators
	

	//Go Conditionals
	//if 
	//if - else 
	//else if 
	//nested if 
	

	//switchs in Go:
	 day := 4

  switch day {
  case 1:
    fmt.Println("Monday")
  case 2:
    fmt.Println("Tuesday")
  case 3:
    fmt.Println("Wednesday")
  case 4:
    fmt.Println("Thursday")
  case 5:
    fmt.Println("Friday")
  case 6:
    fmt.Println("Saturday")
  case 7:
    fmt.Println("Sunday")
  default:
	fmt.Println("this is the default")
  }

  //mutli-case switch:
	newday := 7

   switch newday {
   case 1,3,5:
    fmt.Println("Odd weekday")
   case 2,4:
     fmt.Println("Even weekday")
   case 6,7:
    fmt.Println("Weekend")
  default:
    fmt.Println("Invalid day of day number")
  }

  //loops in Go:
  //everything is common aside new 'range' keyword:
  justSomething := []int{3,4,5,6,7,8,9,0} 
  for i, val := range justSomething {
	fmt.Println(i, val)
  }

  firstFunction("just a paramenter");

  var a int = th();
  fmt.Println(a);


   //Go structures:
  //A struct (short for structure) is used to create a collection of members of different data types, into a single variable.
 //While arrays are used to store multiple values of the same data type into a single variable, structs are used to store multiple values of different data types into a single variable.
//A struct can be useful for grouping data together to create records.

type Person struct {
  name string
  age int
  job string
  salary int
}

  var pers1 Person
  var pers2 Person

  // Pers1 specification
  pers1.name = "Hege"
  pers1.age = 45
  pers1.job = "Teacher"
  pers1.salary = 6000

  // Pers2 specification
  pers2.name = "Cecilie"
  pers2.age = 24
  pers2.job = "Marketing"
  pers2.salary = 4500

  // Access and print Pers1 info
  fmt.Println("Name: ", pers1.name)
  fmt.Println("Age: ", pers1.age)
  fmt.Println("Job: ", pers1.job)
  fmt.Println("Salary: ", pers1.salary)

  // Access and print Pers2 info
  fmt.Println("Name: ", pers2.name)
  fmt.Println("Age: ", pers2.age)
  fmt.Println("Job: ", pers2.job)
  fmt.Println("Salary: ", pers2.salary)

  //Go Maps:
}



//go mod init example.com/hello
//go run .\helloworld.go     