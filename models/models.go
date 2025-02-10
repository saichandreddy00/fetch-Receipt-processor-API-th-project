package models

// Receipt structure for the incoming JSON
type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

// Item structure for the items in the receipt
type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

// Response for process receipt
type ProcessResponse struct {
	ID string `json:"id"`
}

// Response for points calculation
type PointsResponse struct {
	Points int `json:"points"`
}
