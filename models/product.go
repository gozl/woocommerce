package models

import (
	"encoding/json"
	"strings"
)

// ProductItem represents a product associated with an order
type ProductItem struct {
	//*TaxableCost
	ID int                 `json:"id,omitempty"`
	Name string            `json:"name,omitempty"`
	ProductID int          `json:"product_id,omitempty"`
	VariationID int        `json:"variation_id,omitempty"`
	Quantity int           `json:"quantity,omitempty"`
	SKU string             `json:"sku,omitempty"`
	TaxClass string        `json:"tax_class,omitempty"`
	Metadata []MetaItem    `json:"meta_data,omitempty"`
	Subtotal string        `json:"subtotal,omitempty"` //fd
	SubtotalTax string     `json:"subtotal_tax,omitempty"` //fd
	Price string           `json:"price,omitempty"` //fd
}

// UnmarshalJSON parses JSON encoded data to a ProductItem struct
func (c *ProductItem) UnmarshalJSON(data []byte) error {
	if data == nil || len(data) == 0 {
		return nil
	}

	type rObj ProductItem
	aux := &struct{
		*rObj
	}{
		rObj: (*rObj)(c),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		if strings.Contains(err.Error(), "cannot unmarshal string into Go struct") && strings.Contains(err.Error(), "[]models.MetaItem") {
			aux2 := &struct{
				Metadata        string `json:"meta_data,omitempty"`
				*rObj
			}{
				rObj: (*rObj)(c),
			}

			if err2 := json.Unmarshal(data, &aux2); err2 != nil {
				return err2
			}
		} else {
			return err
		}
	}

	return nil
}
