package models

import (
	"encoding/json"
)

// TaxableCost is any cost that is taxable
type TaxableCost struct {
	Total float64          `json:"total,omitempty"`
	TotalTax float64       `json:"total_tax,omitempty"`
	Taxes []TaxItem        `json:"taxes,omitempty"`
}

// UnmarshalJSON parses JSON encoded data to a TaxableCost struct
func (c *TaxableCost) UnmarshalJSON(data []byte) error {
	if data == nil || len(data) == 0 {
		return nil
	}

	var objmap map[string]json.RawMessage
	err := json.Unmarshal(data, &objmap)
	if err != nil {
		return err
	}

	_, ok := objmap["total"]
	if ok {
		err = json.Unmarshal(objmap["total"], &c.Total)
		if err != nil {
			c.Total, err = tryMarshalStringAsFloat64(objmap["total"])
			if err != nil {
				return err
			}
		}
	}

	_, ok = objmap["total_tax"]
	if ok {
		err = json.Unmarshal(objmap["total_tax"], &c.TotalTax)
		if err != nil {
			c.TotalTax, err = tryMarshalStringAsFloat64(objmap["total_tax"])
			if err != nil {
				return err
			}
		}
	}

	_, ok = objmap["taxes"]
	if ok {
		err := json.Unmarshal(objmap["taxes"], &c.Taxes)
		if err != nil {
			return err
		}
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

// UnmarshalJSON parses JSON encoded data to a TaxItem struct
func (c *TaxItem) UnmarshalJSON(data []byte) error {
	if data == nil || len(data) == 0 {
		return nil
	}

	var objmap map[string]json.RawMessage
	err := json.Unmarshal(data, &objmap)
	if err != nil {
		return err
	}

	_, ok := objmap["id"]
	if ok {
		err = json.Unmarshal(objmap["id"], &c.ID)
		if err != nil {
			return err
		}
	}

	_, ok = objmap["rate_code"]
	if ok {
		err = json.Unmarshal(objmap["rate_code"], &c.RateCode)
		if err != nil {
			return err
		}
	}

	_, ok = objmap["rate_id"]
	if ok {
		err = json.Unmarshal(objmap["rate_id"], &c.RateID)
		if err != nil {
			return err
		}
	}

	_, ok = objmap["label"]
	if ok {
		err = json.Unmarshal(objmap["label"], &c.Label)
		if err != nil {
			return err
		}
	}

	_, ok = objmap["compound"]
	if ok {
		err = json.Unmarshal(objmap["compound"], &c.Compound)
		if err != nil {
			return err
		}
	}

	_, ok = objmap["total"]
	if ok {
		err = json.Unmarshal(objmap["total"], &c.Total)
		if err != nil {
			c.Total, err = tryMarshalStringAsFloat64(objmap["total"])
			if err != nil {
				return err
			}
		}
	}

	_, ok = objmap["shipping"]
	if ok {
		err = json.Unmarshal(objmap["shipping"], &c.Shipping)
		if err != nil {
			c.Shipping, err = tryMarshalStringAsFloat64(objmap["shipping"])
			if err != nil {
				return err
			}
		}
	}

	_, ok = objmap["rate_percent"]
	if ok {
		err = json.Unmarshal(objmap["rate_percent"], &c.RatePercent)
		if err != nil {
			c.RatePercent, err = tryMarshalStringAsFloat64(objmap["rate_percent"])
			if err != nil {
				return err
			}
		}
	}

	_, ok = objmap["meta_data"]
	if ok {
		c.Metadata, err = tryMarshalMetadata(objmap["meta_data"])
		if err != nil {
			return err
		}
	}

	return nil
}
