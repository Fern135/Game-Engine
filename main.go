package Go_Engine

/*the libraries */
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//---------------//

func main() {

	reader := bufio.NewReader(os.Stdin) //<------- declaring the user input

	while true{
		strings.Title("Go Engine")//<-------------- i think this is the title. i have not compiled this yet. 

	// the welcome screen and menu
	fmt.Println("Welcome !!")
	fmt.Println("A: Start a new project")
	fmt.Println("B: Options")
	fmt.Println("C: Graphics Modifier")
	fmt.Println("D: Close")
	text, _ := reader.ReadString('\n')

	inputVal(text) //<---------------------------- passing the information to the function
	}
}

// validating the user input
func inputVal(input string) {

	// user input validation 
	if input == "A" || input == "a" {
		start()
	} else if input == "B" || input == "b" {
		userOptions()
	} else if input == "C" || input == "c" {
		graphicsMod()
	} else if input == "D" || input == "d" {
		exit()
	} else {
		fmt.Println("Not one of the options")
	}

}

// start the main app for loading the tools
func start() {
	fmt.Println("This is for the start menu ")
	tmp()
}

// closing the menu screen
func exit() {

}

// the graphics usage modifier of the entire game engine
func graphicsMod() {
	tmp()
}

// the options for the modifying in the entire game engine
func userOptions() {
	tmp()
}

// checking if any new projects exist
func checkIfProjectExist() {
	tmp()
}

// temporary hold for the incoming functions. will try to not use OOP 
func tmp() {
	fmt.Println("Temp")
}
