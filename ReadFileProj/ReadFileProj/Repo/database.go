package Repo

import "ReadFileProj/ice_cream"

//struct that stores all the slices of icecream for unmarshaling from json

type DbStruct struct {
	Icecream []ice_cream.IceCream `json:"icecream"`
}
