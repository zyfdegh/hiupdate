package entity

type Report struct {
	Title   string   `json:"title"`
	Date    string   `json:"date"`
	Weekday string   `json:"weekday"`
	Records []Record `json:"records"`
}
