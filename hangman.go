package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func NOTContains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return false
		}
	}
	return true
}

func main() {
	play := 1
	for play > 0 {
		game_list := [3]string{"dell", "apple", "hp"}
		chances := 10
		error_made := 0
		ind := rand.Intn(len(game_list))
		the_word := game_list[ind]

		l := make([]string, 0)

		var letter_count string
		for i := 0; i < len(the_word); i++ {
			l = append(l, "_")

		}
		fmt.Print("--------------------------------------------\n")

		fmt.Print("Guess the Word with ", len(the_word), "no. of letters\n")

		for error_made <= chances {

			if NOTContains(l, "_") {
				fmt.Print("You Won! It was ",the_word)
				fmt.Print("\n--------------------------------------------\n")
				fmt.Print("Want to play again ? PRESS 1 ELSE 0:\n")
				// var Letter string
				fmt.Scanln(&play, "\n")
				if play == 1 {
					main()
				} else {
					break
				}
			} else {
				fmt.Print("Enter the Letter : ")
				var Letter string
				fmt.Scanln(&Letter, "\n")

				if strings.Contains(letter_count, Letter) {
					fmt.Print("Letter Tried!\n") //, letter_count)
				} else if len(Letter) > 1 {
					fmt.Print("\nEnter a word not a Letter!\n")
				} else if strings.Contains(the_word, Letter) {

					fmt.Print(strings.Index(the_word, Letter), "\n")
					l[strings.Index(the_word, Letter)] = Letter
					fmt.Printf("Remaining chances : %v \n", chances-error_made)
					fmt.Print(l, "\n")
					letter_count = letter_count + Letter
				} else {
					fmt.Print("Wrong Guess.", "\n")
					letter_count = letter_count + Letter
					error_made++
					fmt.Printf("Remaining chances : %v \n", chances-error_made)
					if error_made == 1 {
						fmt.Print("|\n")
					} else if error_made == 2 {
						fmt.Print(" \U0001F913\n")
						fmt.Print(" |\n")
					} else if error_made == 3 {
						fmt.Print(" \U0001F913\n")
						fmt.Print(" |\n")
						fmt.Print("/ \\\n")
					} else if error_made == 4 {
						fmt.Print(" \U0001F913\n")
						fmt.Print(" |~\n")
						fmt.Print("/ \\\n")
					} else if error_made == 5 {
						fmt.Print(" \U0001F913\n")
						fmt.Print("~|~\n")
						fmt.Print("/ \\\n")
					} else if error_made == 6 {
						fmt.Print(" |\n")
						fmt.Print(" \U0001F605\n")
						fmt.Print("~|~\n")
						fmt.Print("/ \\\n")
					} else if error_made == 7 {
						fmt.Print("____\n")
						fmt.Print("   |\n")
						fmt.Print("   \U0001F630\n")
						fmt.Print("  ~|~\n")
						fmt.Print("  / \\\n")
					} else if error_made == 8 {
						fmt.Print("____\n")
						fmt.Print("|  |\n")
						fmt.Print("|  \U0001F630\n")
						fmt.Print("| ~|~\n")
						fmt.Print("| / \\\n")
					} else if error_made == 9 {
						fmt.Print("  ____\n")
						fmt.Print("  |  |\n")
						fmt.Print("  |  \U0001F630\n")
						fmt.Print("  | ~|~\n")
						fmt.Print("  | / \\\n")
						fmt.Print("_____________\n")

					} else if error_made == 10 {
						fmt.Printf("You lost! \nThe Word was  : %v \n", the_word)
						fmt.Print("GAME OVER..!!\n")
						fmt.Print("  ____\n")
						fmt.Print("  |  |\n")
						fmt.Print("  |  \U0001F974           YOU HAVE BEEN HANGED  '\U0001F480'\n")
						fmt.Print("  | ~|~\n")
						fmt.Print("  | / \\\n")
						fmt.Print("__|__________\n")
						fmt.Print("--------------------------------------------\n")
						fmt.Print("Want to play again ? PRESS 1 ELSE 0:")
						// var Letter string
						fmt.Scanln(&play, "\n")
						if play == 1 {
							main()
						} else {
							break
						}

					}

				}
				fmt.Print("--------------------------------------------\n")
			}

		}

	}
}
