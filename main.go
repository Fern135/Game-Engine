package Go_Engine

/*the libraries */
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
//---------------//

func main(){

	strings.Title("Go Engine")//<-------------- i think this is the title
	reader := bufio.NewReader(os.Stdin)//<------- declaring the user input

	// the welcome screen and menu
	fmt.Println("Welcome !!")
	fmt.Println("A")
	text, _ := reader.ReadString('\n')

	inputVal(text)//<---------------------------- passing the information to the function

}


// validating the user input
func inputVal(input string){

	if input == "A" || input == "a" {
		start()
	}

}

func start(){
	fmt.Println("This is for the start menu ")
}