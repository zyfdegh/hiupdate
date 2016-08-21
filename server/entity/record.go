package entity

type Record struct {
	ID      string  `json:"id"`
	Date    string  `json:"date"`
	Person  Person  `json:"person"`
	Content Content `json:"content"`
}

type Person struct {
	Name string `json:"name"`
	Team string `json:"team"`
}

type Content struct {
	Done  string `json:"done"`
	Todo  string `json:"todo"`
	Issue string `json:"issue"`
}
