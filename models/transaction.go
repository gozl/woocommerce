package models

// PaymentTransaction represents a transaction that is a payment
type PaymentTransaction struct {
	PayMethod string       `json:"payment_method,omitempty"`
	PayMethodTitle string  `json:"payment_method_title,omitempty"`
	TransactionID string   `json:"transaction_id,omitempty"`
	Paid Timestamp         `json:"date_paid_gmt,omitempty"`
}
