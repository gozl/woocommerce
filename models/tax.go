package models

import (
	"strconv"
	"encoding/json"
	"strings"
)

// TaxableCost is any cost that is taxable
type TaxableCost struct {
	Total float64          `json:"total,omitempty"`
	TotalTax float64       `json:"total_tax,omitempty"`
	Taxes []TaxItem        `json:"taxes,omitempty"`
}

// MarshalJSON serializes a TaxableCost struct to JSON encoded data
func (c *TaxableCost) MarshalJSON() ([]byte, error) {
	type jsonObj TaxableCost
	return json.Marshal(&struct{
		Total        string `json:"total,omitempty"`
		TotalTax     string `json:"total_tax,omitempty"`
		*jsonObj
	}{
		Total: strconv.FormatFloat(c.Total, 'f', -1, 64),
		TotalTax: strconv.FormatFloat(c.TotalTax, 'f', -1, 64),
		jsonObj:  (*jsonObj)(c),
	})
}

// UnmarshalJSON parses JSON encoded data to a TaxableCost struct
func (c *TaxableCost) UnmarshalJSON(data []byte) error {
	if data == nil || len(data) == 0 {
		return nil
	}

	type rObj TaxableCost
	aux := &struct{
		Total        string `json:"total,omitempty"`
		TotalTax     string `json:"total_tax,omitempty"`
		*rObj
	}{
		rObj: (*rObj)(c),
	}

	var totalStr, totalTaxStr string
	if err := json.Unmarshal(data, &aux); err != nil {
		if strings.Contains(err.Error(), "cannot unmarshal string into Go struct field") {
			aux2 := &struct{
				Total        string `json:"total,omitempty"`
				TotalTax     string `json:"total_tax,omitempty"`
				Taxes        string `json:"taxes,omitempty"`
				*rObj
			}{
				rObj: (*rObj)(c),
			}

			if err2 := json.Unmarshal(data, &aux2); err2 != nil {
				return err2
			}

			totalStr = aux2.Total
			totalTaxStr = aux2.TotalTax
		} else {
			return err
		}
	} else {
		totalStr = aux.Total
		totalTaxStr = aux.TotalTax
	}

	var errFloat error

	c.Total, errFloat = strconv.ParseFloat(totalStr, 64)
	if errFloat != nil {
		return errFloat
	}

	c.TotalTax, errFloat = strconv.ParseFloat(totalTaxStr, 64)
	if errFloat != nil {
		return errFloat
	}

	return nil
}

// TaxItem represents a type of tax
type TaxItem struct {
	ID int                 `json:"id,omitempty"`
	RateCode string        `json:"rate_code,omitempty"`
	RateID int             `json:"rate_id,omitempty"`
	Label string           `json:"label,omitempty"`
	Compound bool          `json:"compound,omitempty"`
	Total float64          `json:"tax_total,omitempty"`
	Shipping float64       `json:"shipping_tax_total,omitempty"`
	RatePercent float64    `json:"rate_percent,omitempty"`
	Metadata []MetaItem    `json:"meta_data,omitempty"`
}

// MarshalJSON serializes a TaxItem struct to JSON encoded data
func (c *TaxItem) MarshalJSON() ([]byte, error) {
	type jsonObj TaxItem
	return json.Marshal(&struct{
		Total        string `json:"tax_total,omitempty"`
		Shipping     string `json:"shipping_tax_total,omitempty"`
		*jsonObj
	}{
		Total: strconv.FormatFloat(c.Total, 'f', -1, 64),
		Shipping: strconv.FormatFloat(c.Shipping, 'f', -1, 64),
		jsonObj:  (*jsonObj)(c),
	})
}

// UnmarshalJSON parses JSON encoded data to a TaxItem struct
func (c *TaxItem) UnmarshalJSON(data []byte) error {
	if data == nil || len(data) == 0 {
		return nil
	}

	type rObj TaxItem
	aux := &struct{
		Total        string `json:"tax_total,omitempty"`
		Shipping     string `json:"shipping_tax_total,omitempty"`
		*rObj
	}{
		rObj: (*rObj)(c),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	var errFloat error

	c.Total, errFloat = strconv.ParseFloat(aux.Total, 64)
	if errFloat != nil {
		return errFloat
	}

	c.Shipping, errFloat = strconv.ParseFloat(aux.Shipping, 64)
	if errFloat != nil {
		return errFloat
	}

	return nil
}
