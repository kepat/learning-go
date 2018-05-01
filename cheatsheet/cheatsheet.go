/**
 * Learned from https://www.youtube.com/watch?v=CF9S4QZuV30
 * Teacher - Derek Banas
 * Cheat Sheet - http://www.newthinktank.com/2015/02/go-programming-tutorial/
 */
package main

import "fmt"

func main() {

	// Type godoc fmt 'package' 'function'
	// - To know more about the package or function
	
	fmt.Println("Hello World")

	// Variable cannot change their type
	// Variable Declaration

	var age int = 40
	var favNum float64 = 1.6180339
	randNum := 1

	fmt.Println(age, " ", favNum, randNum)

	const pi float64 = 3.14159265

	var (
		varA = 2
		varB = 3
	)

	fmt.Println (varA, varB)

	var myName string = "Kevin Tan"

	fmt.Println(len(myName), myName + " - So cute eh! \n")

	var isOver40 bool = true

	// Formatting
	fmt.Printf("%.3f \n", pi) // Format
	fmt.Printf("%s \n", "String") // String
	fmt.Printf("%T \n", isOver40) // Data Type
	fmt.Printf("%t \n", isOver40) // Print Boolean
	fmt.Printf("%d \n", 100) // Display Integer
	fmt.Printf("%b \n", 100) // Binary Representation
	fmt.Printf("%c \n", 44) // Character Code
	fmt.Printf("%x \n", 17) // Hexcode
	fmt.Printf("%e \n", 100) // Scientific NOtation

	// Operators
	fmt.Println("true && false = ", true && false)
	fmt.Println("true || false = ", true || false)
	fmt.Println("!true && false = ", !true)

	// Foor Loop
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// Relational Operator
	// == != < > <= >=

	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}

	// Condition
	yourAge := 18

	if yourAge >= 16 {
		fmt.Println("You can drive")
	} else {
		fmt.Println ("You can't drive")
	}

	// Switch
	switch yourAge {
		case 16 : fmt.Println("Go Drive")
		case 18 : fmt.Println("Go Vote")
		default : fmt.Println("Go Have Fun")
	}

	// Array
	var favLetters[5] string

	favLetters[0] = "A"
	favLetters[1] = "S"
	favLetters[2] = "P"
	favLetters[3] = "G"
	favLetters[4] = "Z"

	fmt.Println(favLetters[0]);

	favNames := [2]string {"Kevin", "Tan"}
	fmt.Println(favNames[1]);

	for index, value := range favNames {
		fmt.Println(value, index)
	}

	// Slice
	numSlice := []int {5, 4, 3, 2, 1}
	newSlice := numSlice[3: 5]
	leftSlice := numSlice[:2]
	rightSlice := numSlice[2:]

	fmt.Println(newSlice)
	fmt.Println(leftSlice)
	fmt.Println(rightSlice)

	dynamicSlice := make([]int, 5, 10)
	copy(dynamicSlice, numSlice)
	dynamicSlice = append(dynamicSlice, 0, -1)
	fmt.Println(dynamicSlice)
}