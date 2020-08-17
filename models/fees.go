package models

// FeeItem represents fee associated with an order
type FeeItem struct {
	*TaxableCost
	ID int                 `json:"id,omitempty"`
	Name string            `json:"name,omitempty"`
	TaxClass string        `json:"tax_class,omitempty"`
	TaxStatus string       `json:"tax_status,omitempty"`
	Metadata []MetaItem    `json:"meta_data,omitempty"`
}
