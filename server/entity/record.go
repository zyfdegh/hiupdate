package entity

type Record struct {
	Number  int     `json:"number"`
	Person  Person  `json:"person"`
	Content Content `json:"content"`
}

type Person struct {
	Name string `json:"name"`
	Team string `json:"team"`
}

type Content struct {
	Done    string `json:"done"`
	Todo    string `json:"todo"`
	Problem string `json:"problem"`
}
