package main

type Item struct {
	ShortDescription string
	Price            string
}

type Receipt struct {
	Retailer     string
	PurchaseDate string
	PurchaseTime string
	Items        []Item
	Total        string
}

type ReceiptId struct {
	Id string `json:"id"`
}

type Points struct {
	Points int `json:"points"`
}
