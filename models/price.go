package models

import (
	"encoding/json"
)

// Price represents an order's price
type Price struct {
	Currency string        `json:"currency,omitempty"`
	CurrencySymbol string  `json:"currency_symbol,omitempty"`
	Discount float64       `json:"discount_total,omitempty"`
	DiscountTax float64    `json:"discount_tax,omitempty"`
	ShippingCost float64   `json:"shipping_total,omitempty"`
	ShippingTax float64    `json:"shipping_tax,omitempty"`
	CartTax float64        `json:"cart_tax,omitempty"`
	Total float64          `json:"total,omitempty"`
	TotalTax float64       `json:"total_tax,omitempty"`
	TaxIncluded bool       `json:"prices_include_tax,omitempty"`
}

// UnmarshalJSON parses JSON encoded data to a Price struct
func (c *Price) UnmarshalJSON(data []byte) error {
	if data == nil || len(data) == 0 {
		return nil
	}

	var objmap map[string]json.RawMessage
	err := json.Unmarshal(data, &objmap)
	if err != nil {
		return err
	}

	_, ok := objmap["currency"]
	if ok {
		err = json.Unmarshal(objmap["currency"], &c.Currency)
		if err != nil {
			return err
		}
	}

	_, ok = objmap["currency_symbol"]
	if ok {
		err = json.Unmarshal(objmap["currency_symbol"], &c.CurrencySymbol)
		if err != nil {
			return err
		}
	}

	_, ok = objmap["prices_include_tax"]
	if ok {
		err = json.Unmarshal(objmap["prices_include_tax"], &c.TaxIncluded)
		if err != nil {
			return err
		}
	}

	_, ok = objmap["discount_total"]
	if ok {
		err = json.Unmarshal(objmap["discount_total"], &c.Discount)
		if err != nil {
			c.Discount, err = tryMarshalStringAsFloat64(objmap["discount_total"])
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

	_, ok = objmap["shipping_total"]
	if ok {
		err = json.Unmarshal(objmap["shipping_total"], &c.ShippingCost)
		if err != nil {
			c.ShippingCost, err = tryMarshalStringAsFloat64(objmap["shipping_total"])
			if err != nil {
				return err
			}
		}
	}

	_, ok = objmap["shipping_tax"]
	if ok {
		err = json.Unmarshal(objmap["shipping_tax"], &c.ShippingTax)
		if err != nil {
			c.ShippingTax, err = tryMarshalStringAsFloat64(objmap["shipping_tax"])
			if err != nil {
				return err
			}
		}
	}

	_, ok = objmap["cart_tax"]
	if ok {
		err = json.Unmarshal(objmap["cart_tax"], &c.CartTax)
		if err != nil {
			c.CartTax, err = tryMarshalStringAsFloat64(objmap["cart_tax"])
			if err != nil {
				return err
			}
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

	return nil
}