package model

import "github.com/google/uuid"

type Checkout struct {
	ID    uuid.UUID      `json:"id"`
	Items []CheckoutItem `json:"items"`
}
