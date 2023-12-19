package main

import (
	"fmt"
	"goproject/utils"
	"os"
)

func main() {
	var answer int
	var err error
	// Call functions from other packages
	//show the questions
	fmt.Println("select one :")

	question1, question2 := utils.GetRandomQuestions(2)

	fmt.Printf("Q1. %s\n", question1)
	fmt.Printf("Q2. %s\n", question2)

	//take the user-input call in function
	_, err = fmt.Scanf("%d", &answer)

	if err != nil {
		fmt.Println("error occured : ", err)
		return
	} else {
		if answer == 1 || answer == 2 {
			//call the function to retrieve the question number index search the index in solution variable
			solution := utils.GetSolution(answer)
			fmt.Printf("Yudhistira : %s\n ", solution)
		} else {
			fmt.Println("Invalid Input")
		}
	}
	fmt.Println("Thank you :)")
	os.Exit(0)
}
