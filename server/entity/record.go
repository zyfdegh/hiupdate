package entity

// Record is the structure of one's daily update
type Record struct {
	ID      string  `json:"id"`
	Date    string  `json:"date"`
	Person  Person  `json:"person"`
	Content Content `json:"content"`
}

// Person is the structure of a employer
type Person struct {
	Name string `json:"name"`
	Team string `json:"team"`
}

// Content is the data written
type Content struct {
	Done  string `json:"done"`
	Todo  string `json:"todo"`
	Issue string `json:"issue"`
}
