package Repo

import (
	"ReadFileProj/ice_cream"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Create a function that accepts an IceCream, converts it to JSON and stores it on the next available line in iceCreamDb.txt
// Return an error if any of the functions you call return an error
func CreateIceCream(ice ice_cream.IceCream) error {

	file, _ := ioutil.ReadFile("Repo/iceCreamDB.json") //read into file

	dbMap := map[string][]interface{}{}
	err := json.Unmarshal([]byte(file), &dbMap) //unmarshal into struct
	if err != nil {
		fmt.Println(err)
	}

	dbMap["icecream"] = append(dbMap["icecream"], ice)

	Marshaled, err := json.Marshal(&dbMap)

	ioutil.WriteFile("Repo/iceCreamDB.json", Marshaled, 0644)

	return err

}

// Create function that can retrieve all the ice creams from iceCreamDb.txt
// The function should unmarshal each one into an ice cream struct and return a slice of IceCream --> []IceCream
// Return an error if any of the functions you call return an error

func GetAllIceCreams() ([]ice_cream.IceCream, error) {
	file, _ := ioutil.ReadFile("Repo/iceCreamDB.json")
	p := DbStruct{}

	err := json.Unmarshal([]byte(file), &p)

	if err != nil {
		fmt.Println(err)
	}

	return p.Icecream, err

}

//check the db if that icecream exists
func GetIceCreamByName(name string) (ice_cream.IceCream, bool) {
	file, _ := ioutil.ReadFile("Repo/iceCreamDB.json")
	p := DbStruct{}

	err := json.Unmarshal([]byte(file), &p)

	if err != nil {
		fmt.Println(err)
	}
	for _, v := range p.Icecream {
		if v.Name == name {
			return v, true
		}
	}
	return ice_cream.IceCream{}, false

}

func DeleteByName(name string) (bool, []ice_cream.IceCream) {
	file, _ := ioutil.ReadFile("Repo/iceCreamDB.json")
	p := DbStruct{}

	err := json.Unmarshal([]byte(file), &p)

	if err != nil {
		fmt.Println(err)
	}

	for i, v := range p.Icecream {
		if v.Name == name {
			p.Icecream = append(p.Icecream[:i], p.Icecream[i+1:]...)
			Marshaled, _ := json.Marshal(&p)

			ioutil.WriteFile("Repo/iceCreamDB.json", Marshaled, 0644)

			return true, p.Icecream
		}
	}
	return false, p.Icecream

}
