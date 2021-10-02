package Input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//This file takes input from the user to be returned to the user
func UserInputToGet() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Pls, pick an icecream to return:")
	usrInput, _ := reader.ReadString('\n')
	word := strings.Replace(usrInput, "\n", "", -1)

	return word
}

//This file takes input to be deleted from the DB
func UserInputToDelete() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Pls, pick an icecream to delete:")
	usrInput, _ := reader.ReadString('\n')
	word := strings.Replace(usrInput, "\n", "", -1)

	return word
}
