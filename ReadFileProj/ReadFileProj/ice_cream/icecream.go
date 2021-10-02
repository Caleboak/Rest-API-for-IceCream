package ice_cream

// Add the appropriate fields and JSON tags to the struct below
type IceCream struct {
	Name     string   `json:"Name"`
	Flavour  string   `json:"Flavor"`
	Base     string   `json:"Base"`
	Toppings []string `json:"Toppings"`
}
