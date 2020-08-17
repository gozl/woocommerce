package models

import (
	"strconv"
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
	Subtotal float64       `json:"subtotal,omitempty"` //fd
	SubtotalTax float64    `json:"subtotal_tax,omitempty"` //fd
	Price float64          `json:"price,omitempty"` //fd
}

// MarshalJSON serializes a ProductItem struct to JSON encoded data
func (c *ProductItem) MarshalJSON() ([]byte, error) {
	type jsonObj ProductItem
	return json.Marshal(&struct{
		Price           string `json:"price,omitempty"`
		Subtotal        string `json:"subtotal,omitempty"`
		SubtotalTax     string `json:"subtotal_tax,omitempty"`
		*jsonObj
	}{
		Price: strconv.FormatFloat(c.Price, 'f', -1, 64),
		Subtotal: strconv.FormatFloat(c.Subtotal, 'f', -1, 64),
		SubtotalTax: strconv.FormatFloat(c.SubtotalTax, 'f', -1, 64),
		jsonObj:  (*jsonObj)(c),
	})
}

// UnmarshalJSON parses JSON encoded data to a ProductItem struct
func (c *ProductItem) UnmarshalJSON(data []byte) error {
	if data == nil || len(data) == 0 {
		return nil
	}

	type rObj ProductItem
	aux := &struct{
		Subtotal        string `json:"subtotal,omitempty"`
		SubtotalTax     string `json:"subtotal_tax,omitempty"`
		*rObj
	}{
		rObj: (*rObj)(c),
	}
	//aux.TaxableCost = &TaxableCost{}

	var subtotalStr, subtotalTaxStr string
	if err := json.Unmarshal(data, &aux); err != nil {
		if strings.Contains(err.Error(), "cannot unmarshal string into Go struct") && strings.Contains(err.Error(), "[]models.MetaItem") {
			aux2 := &struct{
				Metadata        string `json:"meta_data,omitempty"`
				Subtotal        string `json:"subtotal,omitempty"`
				SubtotalTax     string `json:"subtotal_tax,omitempty"`
				*rObj
			}{
				rObj: (*rObj)(c),
			}

			if err2 := json.Unmarshal(data, &aux2); err2 != nil {
				return err2
			}

			subtotalStr = aux2.Subtotal
			subtotalTaxStr = aux2.SubtotalTax
		} else {
			return err
		}
	} else {
		subtotalStr = aux.Subtotal
		subtotalTaxStr = aux.SubtotalTax
	}

	var errFloat error

	/*
	c.Price, errFloat = strconv.ParseFloat(aux.Price, 64)
	if errFloat != nil {
		return errFloat
	}
	*/

	c.Subtotal, errFloat = strconv.ParseFloat(subtotalStr, 64)
	if errFloat != nil {
		return errFloat
	}
	c.SubtotalTax, errFloat = strconv.ParseFloat(subtotalTaxStr, 64)
	if errFloat != nil {
		return errFloat
	}

	return nil
}
