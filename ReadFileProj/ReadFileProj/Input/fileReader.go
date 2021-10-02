package Input

import (
	"ReadFileProj/ice_cream"
	"encoding/json"
	"io/ioutil"
)

//Use the appropriate functions from the os package, JSON package, ioutils package etc. To read the input
// from a file, unmarshal it into an IceCream struct and return the struct
// Return an error if any of your functions that you call return an error

func ReadFile(fileName string) (ice_cream.IceCream, error) {

	file, _ := ioutil.ReadFile(fileName) //read into file

	data := ice_cream.IceCream{} //create an icecream instance

	err := json.Unmarshal([]byte(file), &data) //unmarshal into struct

	return data, err //return the "unmarshaled into struct" and err
}
