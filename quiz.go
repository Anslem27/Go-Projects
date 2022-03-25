/* Statically typed language */
/* Unused variables return errors */

/* var name2 string = "Another name"
var number int = 1000
var positive uint = 2
var float float64 = -92.3
name := "Eddie"
age := 12
*/
//! GoLang Quiz game

package main

import "fmt"

func main() {
	fmt.Println("Welcome to my Quiz game")
	var name string
	fmt.Println("Kindly input your name")
	fmt.Scan(&name)
	fmt.Printf("Great name %v ☺", name)
	fmt.Println("\nKindly input your age")
	var age uint
	fmt.Scan(&age)
	if age >= 10 {
		fmt.Printf("OOoh great %v, %v already. Get ready to PLAY ", name, age)
	} else {
		fmt.Println("Sorry the game is too brutal for people this age")
		return
	}

	fmt.Println("\n Lets continue....")

	//user score
	score := 0
	//no of qns
	noOfQuestions := 2

	fmt.Printf("What is better, Samsung or Oneplus? : ")
	var qnAnswer string
	fmt.Scan(&qnAnswer)
	if qnAnswer == "Oneplus" {
		fmt.Println("Correct 😁 ")
		score++
		/* or score = score+1 or score +=1 */
	} else if qnAnswer == "oneplus" {
		score++
		fmt.Println("Correct 😁 ")
	} else {
		fmt.Println("Oooops we all know you couldnt get this one ryt")
	}

	println("Next Qn......")

	fmt.Printf("\nWhat is better, Nvidia or AMD? : ")
	var qnAnswer2 string
	fmt.Scan(&qnAnswer2)
	if qnAnswer2 == "Nvidia" || qnAnswer2 == "nvidia" {
		fmt.Println("Correct 😁 ")
		score++
		/* or score = score+1 or score +=1 */
	} else {
		fmt.Println("Believe it or not")
	}

	//but an int/int =int hence converting to doubles or float
	//intead of using percentage := (score / noOfQuestions) * 100

	//? percentage variable

	percentage := (float64(score) / float64(noOfQuestions)) * 100
	fmt.Printf("You scored %v out of %v qns", score, noOfQuestions)
	fmt.Printf("\n Which is about %v%% ", percentage)
	// logical operators
	// || => or
}
