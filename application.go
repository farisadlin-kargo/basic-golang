package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var name string
	x := os.Args[1]
	s, _ := strconv.ParseInt(x, 10, 0)
	y, err := strconv.Atoi(x)

	if len(os.Args) > 2 {
		fmt.Println("Invalid number of arguments")
	} else {
		if err != nil {
			fmt.Println("strconv.Atoi:", err)
		} else {

			if s < 0 || s > 10 {
				fmt.Println("The input number should be greater than 0 and less than 10")
			} else {

				fmt.Println("Your name please? Press the Enter key when done.")
				in := bufio.NewReader(os.Stdin)
				name, _ = in.ReadString('\n')

				if len(name) > 1 {
					for i := 0; i < y; i++ {
						fmt.Print("Nice to meet you ", name)
					}
				} else {
					fmt.Print("You didn't enter your name")
				}

			}
		}

	}

}
