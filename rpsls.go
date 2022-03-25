package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var scoreUser int
var scorePc int
var scoreTotal int
var scoreTie int
var x int

// Returns an int >= min, < max
func randomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano()) //Crear semilla para que el numero aleatorio no se repita.
	return min + rand.Intn(max-min)
}

func userElection() int {
	reader := bufio.NewReader(os.Stdin) //Create reader
	option, _ := reader.ReadString('\n')
	option = strings.Replace(option, "\n", "", -1)

	var option_final int
	var err error

	if option == "e" {
		welcome()
		fmt.Println("You choose to exit the application. See you next time!")
		if scoreTotal != 0 {
			fmt.Println("-------------------")
			fmt.Println("Scores: ")
			fmt.Println("User:", scoreUser)
			fmt.Println("PC:", scorePc)
			fmt.Println("Tie:", scoreTie)
			fmt.Println("-------------------")
			fmt.Println("Winning percentage:", percentage(scoreTotal, scoreUser), "%")
			fmt.Println("Losing percentage:", percentage(scoreTotal, scorePc), "%")
			fmt.Println("Draw percentage:", percentage(scoreTotal, scoreTie), "%")
			fmt.Println("-------------------")
			os.Exit(0)
		} else {
			os.Exit(0)
		}

	} else if option != "e" {
		option_final, err = strconv.Atoi(option)
		if err != nil {
			option_final = 0
		}
	}
	return option_final
}

// Select the winner based on the User and PC selection
func decideWinner(pcSelection, userSelection string) {
	//------ User win------
	if pcSelection == "rock" && userSelection == "paper" {
		fmt.Println("You won, paper wraps rock!")
		userWin := 1
		scoreUser = scoreUser + userWin
		scoreTotal = scoreUser + scorePc + scoreTie
	} else if pcSelection == "paper" && userSelection == "scissor" {
		fmt.Println("You won, scissor cuts paper!")
		userWin := 1
		scoreUser = scoreUser + userWin
		scoreTotal = scoreUser + scorePc + scoreTie
	} else if pcSelection == "scissor" && userSelection == "rock" {
		fmt.Println("You won, rock crushes scissor!")
		userWin := 1
		scoreUser = scoreUser + userWin
		scoreTotal = scoreUser + scorePc + scoreTie
	} else if pcSelection == "lizard" && userSelection == "rock" {
		fmt.Println("You won, rock crushes lizard!")
		userWin := 1
		scoreUser = scoreUser + userWin
		scoreTotal = scoreUser + scorePc + scoreTie
	} else if pcSelection == "lizard" && userSelection == "scissor" {
		fmt.Println("You won, scissor decapites lizard!")
		userWin := 1
		scoreUser = scoreUser + userWin
		scoreTotal = scoreUser + scorePc + scoreTie
	} else if pcSelection == "scissor" && userSelection == "spock" {
		fmt.Println("You won, spock breaks scissor!")
		userWin := 1
		scoreUser = scoreUser + userWin
		scoreTotal = scoreUser + scorePc + scoreTie
	} else if pcSelection == "spock" && userSelection == "lizard" {
		fmt.Println("You won, lizard poisons spock!")
		userWin := 1
		scoreUser = scoreUser + userWin
		scoreTotal = scoreUser + scorePc + scoreTie
	} else if pcSelection == "paper" && userSelection == "lizard" {
		fmt.Println("You won, lizard eats paper!")
		userWin := 1
		scoreUser = scoreUser + userWin
		scoreTotal = scoreUser + scorePc + scoreTie
	} else if pcSelection == "spock" && userSelection == "paper" {
		fmt.Println("You won, paper disavows spock!")
		userWin := 1
		scoreUser = scoreUser + userWin
		scoreTotal = scoreUser + scorePc + scoreTie
	} else if pcSelection == "rock" && userSelection == "spock" {
		fmt.Println("You won, spock vaporizes rock!")
		userWin := 1
		scoreUser = scoreUser + userWin
		scoreTotal = scoreUser + scorePc + scoreTie
	}

	//------ IA win ------
	if userSelection == "rock" && pcSelection == "paper" {
		fmt.Println("You lost, paper wraps rock!")
		pcWin := 1
		scorePc = scorePc + pcWin
		scoreTotal = scoreUser + scorePc + scoreTie
	} else if userSelection == "paper" && pcSelection == "scissor" {
		fmt.Println("You lost, scissor cuts paper!")
		pcWin := 1
		scorePc = scorePc + pcWin
		scoreTotal = scoreUser + scorePc + scoreTie
	} else if userSelection == "scissor" && pcSelection == "rock" {
		fmt.Println("You lost, rock crushes scissor!")
		pcWin := 1
		scorePc = scorePc + pcWin
		scoreTotal = scoreUser + scorePc + scoreTie
	} else if userSelection == "lizard" && pcSelection == "rock" {
		fmt.Println("You lost, rock crushes lizard!")
		pcWin := 1
		scorePc = scorePc + pcWin
		scoreTotal = scoreUser + scorePc + scoreTie
	} else if userSelection == "lizard" && pcSelection == "scissor" {
		fmt.Println("You lost, scissor decapites lizard!")
		pcWin := 1
		scorePc = scorePc + pcWin
		scoreTotal = scoreUser + scorePc + scoreTie
	} else if userSelection == "scissor" && pcSelection == "spock" {
		fmt.Println("You lost, spock breaks scissor!")
		pcWin := 1
		scorePc = scorePc + pcWin
		scoreTotal = scoreUser + scorePc + scoreTie
	} else if userSelection == "spock" && pcSelection == "lizard" {
		fmt.Println("You lost, lizard poisons spock!")
		pcWin := 1
		scorePc = scorePc + pcWin
		scoreTotal = scoreUser + scorePc + scoreTie
	} else if userSelection == "paper" && pcSelection == "lizard" {
		fmt.Println("You lost, lizard eats paper!")
		pcWin := 1
		scorePc = scorePc + pcWin
		scoreTotal = scoreUser + scorePc + scoreTie
	} else if userSelection == "spock" && pcSelection == "paper" {
		fmt.Println("You lost, paper disavows spock!")
		pcWin := 1
		scorePc = scorePc + pcWin
		scoreTotal = scoreUser + scorePc + scoreTie
	} else if userSelection == "rock" && pcSelection == "spock" {
		fmt.Println("You lost, spock vaporizes rock!")
		pcWin := 1
		scorePc = scorePc + pcWin
		scoreTotal = scoreUser + scorePc + scoreTie

		//------ Tie ------
	} else if pcSelection == userSelection {
		fmt.Println("It's a tie!")
		tie := 1
		scoreTie = scoreTie + tie
		scoreTotal = scoreUser + scorePc + scoreTie
	}

}

func percentage(scoreTotal, scoreUserPC int) float32 {
	x = scoreUserPC * 100 / scoreTotal
	return float32(x)
}

func welcome() {
	fmt.Printf("\x1bc")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("******************************************************")
	fmt.Println("*                                                    *")
	fmt.Println("*   Welcome to Rock, Paper, Scissor, Lizard, Spock   *")
	fmt.Println("*                                                    *")
	fmt.Println("******************************************************")
	fmt.Println("")
}

func main() {

	welcome()
	for true {
		fmt.Println("Please enter a number or 'e' to exit the program ")
		fmt.Println("1) Rock")
		fmt.Println("2) Paper")
		fmt.Println("3) Scissor")
		fmt.Println("4) Lizard")
		fmt.Println("5) Spock")
		fmt.Println("6) *** Show score ***")
		fmt.Println("")
		fmt.Print("Your election: ")

		option_final := userElection()
		for {
			if option_final >= 1 && option_final <= 6 {
				break
			} else {
				fmt.Println("Error, you must enter a number between 1 and 6.", "'", option_final, "'", "it's not a number")
				fmt.Print("Enter the number: ")
				option_final = userElection()
			}
		}
		var userSelection string

		option_final_pc := randomInt(1, 6)
		var pcSelection string

		//User choise
		if option_final == 1 {
			userSelection = "rock"
		} else if option_final == 2 {
			userSelection = "paper"
		} else if option_final == 3 {
			userSelection = "scissor"
		} else if option_final == 4 {
			userSelection = "lizard"
		} else if option_final == 5 {
			userSelection = "spock"
		} else if option_final == 6 {
			welcome()
			fmt.Println("you choose to see the score!")
			fmt.Println("Scores: ")
			fmt.Println("User:", scoreUser)
			fmt.Println("PC:", scorePc)
			fmt.Println("Tie:", scoreTie)
			fmt.Println("-------------------")
			fmt.Println("Winning percentage:", percentage(scoreTotal, scoreUser), "%")
			fmt.Println("Losing percentage:", percentage(scoreTotal, scorePc), "%")
			fmt.Println("Draw percentage:", percentage(scoreTotal, scoreTie), "%")
			fmt.Println("\nPress any button to continue")
			fmt.Scanln() // wait for Enter Key
			//fmt.Println("-------------------")
			welcome()
			continue

		} else {
			println("Wrong choice!")
		}

		//IA Choise
		if option_final_pc == 1 {
			pcSelection = "rock"
		} else if option_final_pc == 2 {
			pcSelection = "paper"
		} else if option_final_pc == 3 {
			pcSelection = "scissor"
		} else if option_final_pc == 4 {
			pcSelection = "lizard"
		} else if option_final_pc == 5 {
			pcSelection = "spock"
		}

		welcome()
		fmt.Println("-----------")
		fmt.Println("The User choice is:", userSelection)
		fmt.Println("The IA choice is:", pcSelection)

		fmt.Println("-----------")
		fmt.Println("")

		decideWinner(pcSelection, userSelection)

		fmt.Println("")
		fmt.Println("-----------")

		var again string

		for again == "" {
			fmt.Println("Do you want to play again? y/n")
			reader := bufio.NewReader(os.Stdin) //Create reader
			again, _ := reader.ReadString('\n')
			again = strings.Replace(again, "\n", "", -1)
			if again == "y" || again == "Y" {
				welcome()
				break
			} else if again == "n" || again == "N" {
				welcome()
				fmt.Println("You choose to exit the application. See you next time!")
				fmt.Println("-------------------")
				fmt.Println("Scores: ")
				fmt.Println("User:", scoreUser)
				fmt.Println("PC:", scorePc)
				fmt.Println("Tie:", scoreTie)
				fmt.Println("-------------------")
				fmt.Println("Winning percentage:", percentage(scoreTotal, scoreUser), "%")
				fmt.Println("Losing percentage:", percentage(scoreTotal, scorePc), "%")
				fmt.Println("Draw percentage:", percentage(scoreTotal, scoreTie), "%")
				fmt.Println("-------------------")
				os.Exit(0)
			} else {
				fmt.Println("invalid input")
				again = ""
				continue
			}
		}
	}
}
