package app

// structure describing the data fields of each car
type Car struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Year   int    `json:"year"`
	Colour string `json:"colour"`
	Price  int    `json:"price"`
	Image  int    `json:"image"`
}
