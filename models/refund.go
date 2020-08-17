package models

import (
	"strconv"
	"encoding/json"
)

// RefundItem represents a refund associated with an order
type RefundItem struct {
	ID int                 `json:"id,omitempty"`
	Reason string          `json:"reason,omitempty"`
	Total float64          `json:"total,omitempty"`
}

// MarshalJSON serializes a RefundItem struct to JSON encoded data
func (c *RefundItem) MarshalJSON() ([]byte, error) {
	type jsonObj RefundItem
	return json.Marshal(&struct{
		Total           string `json:"total,omitempty"`
		*jsonObj
	}{
		Total: strconv.FormatFloat(c.Total, 'f', -1, 64),
		jsonObj:  (*jsonObj)(c),
	})
}

// UnmarshalJSON parses JSON encoded data to a RefundItem struct
func (c *RefundItem) UnmarshalJSON(data []byte) error {
	if data == nil || len(data) == 0 {
		return nil
	}

	type rObj RefundItem
	aux := &struct{
		Total           string `json:"total,omitempty"`
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

	return nil
}
