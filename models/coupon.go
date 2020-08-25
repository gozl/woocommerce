package models

import (
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

// UnmarshalJSON parses JSON encoded data to a CouponItem struct
func (c *CouponItem) UnmarshalJSON(data []byte) error {
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

	_, ok = objmap["code"]
	if ok {
		err = json.Unmarshal(objmap["code"], &c.Code)
		if err != nil {
			return err
		}
	}

	_, ok = objmap["meta_data"]
	if ok {
		c.Metadata, err = tryMarshalMetadata(objmap["meta_data"])
		if err != nil {
			return err
		}
	}

	_, ok = objmap["discount"]
	if ok {
		err = json.Unmarshal(objmap["discount"], &c.Discount)
		if err != nil {
			c.Discount, err = tryMarshalStringAsFloat64(objmap["discount"])
			if err != nil {
				return err
			}
		}
	}

	_, ok = objmap["discount_tax"]
	if ok {
		err = json.Unmarshal(objmap["discount_tax"], &c.DiscountTax)
		if err != nil {
			c.DiscountTax, err = tryMarshalStringAsFloat64(objmap["discount_tax"])
			if err != nil {
				return err
			}
		}
	}

	return nil
}
