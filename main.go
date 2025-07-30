package main 

import (
	"fmt"

	"os"

	"os/user"

	"go-work/repl"
)

func main(){
	// gets the usernmae first
	user, err := user.Current()

	// similar to an except block 
	if err != nil {
		// panic is same as raise for errors
		panic(err)
	}

	// gets user and allows program to run
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)

	fmt.Printf("Feel free to type in commands\n")

	repl.Start(os.Stdin, os.Stdout)

}