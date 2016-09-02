package entity

// Req to /record
type ReqRecord struct {
	Name  string `json:"name"`
	Done  string `json:"done"`
	Todo  string `json:"todo"`
	Issue string `json:"issue"`
}
