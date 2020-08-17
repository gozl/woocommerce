package models

import (
	"strconv"
	"encoding/json"
)

// CouponItem represents a coupon relating to an order
type CouponItem struct {
	ID int                 `json:"id,omitempty"`
	Code string            `json:"code,omitempty"`
	Discount float64       `json:"discount,omitempty"`
	DiscountTax float64    `json:"discount_tax,omitempty"`
	Metadata []MetaItem    `json:"meta_data,omitempty"`
}

// MarshalJSON serializes a CouponItem struct to JSON encoded data
func (c *CouponItem) MarshalJSON() ([]byte, error) {
	type jsonObj CouponItem
	return json.Marshal(&struct{
		Discount        string `json:"discount,omitempty"`
		DiscountTax     string `json:"discount_tax,omitempty"`
		*jsonObj
	}{
		Discount: strconv.FormatFloat(c.Discount, 'f', -1, 64),
		DiscountTax: strconv.FormatFloat(c.DiscountTax, 'f', -1, 64),
		jsonObj:  (*jsonObj)(c),
	})
}

// UnmarshalJSON parses JSON encoded data to a CouponItem struct
func (c *CouponItem) UnmarshalJSON(data []byte) error {
	if data == nil || len(data) == 0 {
		return nil
	}

	type rObj CouponItem
	aux := &struct{
		Discount        string `json:"discount,omitempty"`
		DiscountTax     string `json:"discount_tax,omitempty"`
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

	return nil
}
