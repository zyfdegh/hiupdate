package entity

type Report struct {
	Title   string   `json:"title"`
	Date    string   `json:"date"`
	Week    string   `json:"week"`
	Records []Record `json:"records"`
}
