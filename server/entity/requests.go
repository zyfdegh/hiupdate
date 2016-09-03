package entity

// ReqRecord is the request to PUT /record
type ReqRecord struct {
	Name  string `json:"name"`
	Done  string `json:"done"`
	Todo  string `json:"todo"`
	Issue string `json:"issue"`
}
