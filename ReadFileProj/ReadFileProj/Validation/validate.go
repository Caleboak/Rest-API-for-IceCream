package Validations

import (
	"ReadFileProj/ice_cream"
	"encoding/json"

	"io/ioutil"
)

// Write a function that takes in an ice cream then reads the file toppings.json.
// Iterate through the toppings returned from toppings.json and iterate throught the toppings in the IceCream
// to verify that the IceCream ony contains toppings listed in toppings.json.
// If the ice cream fails validation return false, otherwise return true.

func Validate(ic ice_cream.IceCream) (bool, error) {
	file, _ := ioutil.ReadFile("icInput/toppings.json")
	//anonymous struct to unmarshal the toppings(json) into
	ValidateToppings := struct {
		Toppings []string
	}{}
	err := json.Unmarshal(file, &ValidateToppings)
	if err != nil {
		return false, err
	}

	compareMap := map[string]bool{}

	for _, v := range ValidateToppings.Toppings {
		compareMap[v] = true

	}
	for _, value := range ic.Toppings {
		if _, ok := compareMap[value]; !ok {
			return false, nil
		}
	}
	return true, nil
}
