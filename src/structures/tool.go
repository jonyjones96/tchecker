package structures

// Tool is an item sold by a builders merchant
type Tool struct {
	Name    string   `json:"productName"`
	Price   float64  `json:"productPrice"`
	Code    string   `json:"productCode"`
	Details []string `json:"productFeatures"`
}
