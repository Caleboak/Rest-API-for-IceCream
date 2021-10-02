package main

import (
	"ReadFileProj/Input"
	"ReadFileProj/Repo"
	Validations "ReadFileProj/Validation"
	"fmt"
)

// Import the necessary packages and call the appropriate methods to read a files input and convert it to an IceCream
// You should then validate IceCream if it fails validation print that out.
// If the icecream passes validation, store the IceCream in the db

var slice = []string{"icInput/ic1.json", "icInput/ic2.json",
	"icInput/ic3.json", "icInput/ic4.json"}

func main() {
	for _, v := range slice {
		res, err := Input.ReadFile(v)
		if err != nil {
			fmt.Println(err)
		}

		validate, err := Validations.Validate(res)
		if err != nil {
			fmt.Println(err)
			return
		} else if validate {
			Repo.CreateIceCream(res)

		} else if !validate {
			fmt.Println("Toppings don't match")
		}

	}
	outputIcecream, _ := Repo.GetAllIceCreams()
	//output the content of the DB to user
	fmt.Printf("%+v\n", outputIcecream)

	//Get the icecream from user to be returned
	getEntry := Input.UserInputToGet()

	//pass the userInput to the function to check the DB
	iceCreamOutput, found := Repo.GetIceCreamByName(getEntry)
	if !found {
		fmt.Println("Not in the Database") //icecream is not in the db
	} else {
		fmt.Println(iceCreamOutput)
	}

	deleteEntry := Input.UserInputToDelete()

	found, remaniningSlice := Repo.DeleteByName((deleteEntry))
	if found {
		fmt.Println("Entry deleted, Current content:")
		fmt.Printf("%+v\n", remaniningSlice)
	} else {
		fmt.Println("Entry not found, returned initial content")
		fmt.Printf("%+v\n", remaniningSlice)

	}

	// Read in the input and unmarshal to an IceCream struct

	// Compare the toppings to the toppings in toppings.json

	// If validation fails returns false from the Validate function and return from the main function

	// If validation passes, marshall the icecream into JSON and store it in iceCreamDb.txt

}
