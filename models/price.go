package models

import (
	"strconv"
	"encoding/json"
)

// Price represents an order's price
type Price struct {
	Currency string        `json:"currency,omitempty"`
	CurrencySymbol string  `json:"currency_symbol,omitempty"`
	Discount float64       `json:"discount_total,omitempty"`
	DiscountTax float64    `json:"discount_tax,omitempty"`
	Shipping float64       `json:"shipping_total,omitempty"`
	ShippingTax float64    `json:"shipping_tax,omitempty"`
	CartTax float64        `json:"cart_tax,omitempty"`
	Total float64          `json:"total,omitempty"`
	TotalTax float64       `json:"total_tax,omitempty"`
	TaxIncluded bool       `json:"prices_include_tax,omitempty"`
}

// MarshalJSON serializes a Price struct to JSON encoded data
func (c *Price) MarshalJSON() ([]byte, error) {
	type jsonObj Price
	return json.Marshal(&struct{
		Discount     string `json:"discount_total,omitempty"`
		DiscountTax  string `json:"discount_tax,omitempty"`
		Shipping     string `json:"shipping_total,omitempty"`
		ShippingTax  string `json:"shipping_tax,omitempty"`
		CartTax      string `json:"cart_tax,omitempty"`
		Total        string `json:"total,omitempty"`
		TotalTax     string `json:"total_tax,omitempty"`
		*jsonObj
	}{
		Discount: strconv.FormatFloat(c.Discount, 'f', -1, 64),
		DiscountTax: strconv.FormatFloat(c.DiscountTax, 'f', -1, 64),
		Shipping: strconv.FormatFloat(c.Shipping, 'f', -1, 64),
		ShippingTax: strconv.FormatFloat(c.ShippingTax, 'f', -1, 64),
		CartTax: strconv.FormatFloat(c.CartTax, 'f', -1, 64),
		Total: strconv.FormatFloat(c.Total, 'f', -1, 64),
		TotalTax: strconv.FormatFloat(c.TotalTax, 'f', -1, 64),
		jsonObj:  (*jsonObj)(c),
	})
}

// UnmarshalJSON parses JSON encoded data to a Price struct
func (c *Price) UnmarshalJSON(data []byte) error {
	if data == nil || len(data) == 0 {
		return nil
	}

	type rObj Price
	aux := &struct{
		Discount     string `json:"discount_total,omitempty"`
		DiscountTax  string `json:"discount_tax,omitempty"`
		Shipping     string `json:"shipping_total,omitempty"`
		ShippingTax  string `json:"shipping_tax,omitempty"`
		CartTax      string `json:"cart_tax,omitempty"`
		Total        string `json:"total,omitempty"`
		TotalTax     string `json:"total_tax,omitempty"`
		*rObj
	}{
		rObj: (*rObj)(c),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	var errFloat error

	c.Discount, errFloat = strconv.ParseFloat(aux.Discount, 64)
	if errFloat != nil {
		return errFloat
	}

	c.DiscountTax, errFloat = strconv.ParseFloat(aux.DiscountTax, 64)
	if errFloat != nil {
		return errFloat
	}

	c.Shipping, errFloat = strconv.ParseFloat(aux.Shipping, 64)
	if errFloat != nil {
		return errFloat
	}

	c.ShippingTax, errFloat = strconv.ParseFloat(aux.ShippingTax, 64)
	if errFloat != nil {
		return errFloat
	}

	c.CartTax, errFloat = strconv.ParseFloat(aux.CartTax, 64)
	if errFloat != nil {
		return errFloat
	}

	c.Total, errFloat = strconv.ParseFloat(aux.Total, 64)
	if errFloat != nil {
		return errFloat
	}

	c.TotalTax, errFloat = strconv.ParseFloat(aux.TotalTax, 64)
	if errFloat != nil {
		return errFloat
	}

	return nil
}
