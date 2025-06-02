package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var cmd string                    // Stores the user's command
	var key string                    // Stores the key name entered by the user
	filename := "keysystemstring.txt" // File where the key is saved
	clr := "nothing here"             // Placeholder text for the "clr" command

	fmt.Println("welcome to key manager")

	for {
		fmt.Println("    ")
		fmt.Print("> ") // Prompt for user input
		fmt.Scan(&cmd)  // Read the command from user

		switch cmd {
		case "exit":
			// Exit the program
			return

		case "new":
			// Create a new key and save it to the file
			fmt.Println("What name? only one line no spaces i dont have time")
			fmt.Print("-> ")
			fmt.Scan(&key)

			// Write the key to the file (overwrites existing content)
			err := os.WriteFile(filename, []byte(key), 0644)
			if err != nil {
				fmt.Println("error saving to file")
				return
			}

		case "clr":
			// Clear the contents of the file by writing placeholder text
			err := os.WriteFile(filename, []byte(clr), 0644)
			if err != nil {
				fmt.Println("error clearing file")
				return
			}

		case "check":
			// Read and display contents of the file
			file, err := os.Open(filename)
			if err != nil {
				fmt.Println("Wrong logic or error")
				return
			}
			defer file.Close()

			// Use a scanner to read the file line by line
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()
				fmt.Println(line) // Print each line
			}

			// Check if any error occurred while scanning
			if err := scanner.Err(); err != nil {
				fmt.Println("unknown logic")
			}

		default:
			// Handle unknown commands
			fmt.Println("Unknown logic")
		}
	}
}
