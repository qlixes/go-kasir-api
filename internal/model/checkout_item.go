package model

type CheckoutItem struct {
	Product  Product `json:"product_id"`
	Quantity int     `json:"quantity"`
}
