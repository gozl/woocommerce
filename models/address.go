package models

// Address represents basic address info
type Address struct {
	Street string    `json:"address_1,omitempty"`
	Street2 string   `json:"address_2,omitempty"`
	City string      `json:"city,omitempty"`
	State string     `json:"state,omitempty"`
	Postcode string  `json:"postcode,omitempty"`
	Country string   `json:"country,omitempty"`
}

// BillingInfo represents an order's billing info
type BillingInfo struct {
	*Address
	FirstName string `json:"first_name,omitempty"`
	LastName string  `json:"last_name,omitempty"`
	Company string   `json:"company,omitempty"`
	Email string     `json:"email,omitempty"`
	Phone string     `json:"phone,omitempty"`
}

// ShippingInfo represents an order's shipping info
type ShippingInfo struct {
	*Address
	FirstName string `json:"first_name,omitempty"`
	LastName string  `json:"last_name,omitempty"`
	Company string   `json:"company,omitempty"`
}
