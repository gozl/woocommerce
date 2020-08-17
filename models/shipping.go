package models

// ShippingMethod represents shipping methods relating to an order
type ShippingMethod struct {
	*TaxableCost
	ID int                 `json:"id,omitempty"`
	MethodTitle string     `json:"method_title,omitempty"`
	MethodID string        `json:"method_id,omitempty"`
	Metadata []MetaItem    `json:"meta_data,omitempty"`
}
