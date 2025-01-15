package models

// Receipt represents the receipt data structure.
type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

// Item represents an item in the receipt.
type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}
