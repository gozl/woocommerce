package models

import (
	"encoding/json"
)

// ProductItem represents a product associated with an order
type ProductItem struct {
	*TaxableCost
	ID int                 `json:"id,omitempty"`
	Name string            `json:"name,omitempty"`
	ProductID int          `json:"product_id,omitempty"`
	VariationID int        `json:"variation_id,omitempty"`
	Quantity int           `json:"quantity,omitempty"`
	SKU string             `json:"sku,omitempty"`
	TaxClass string        `json:"tax_class,omitempty"`
	Metadata []MetaItem    `json:"meta_data,omitempty"`
	Subtotal float64       `json:"subtotal,omitempty"`
	SubtotalTax float64    `json:"subtotal_tax,omitempty"`
	Price float64          `json:"price,omitempty"`
}

// UnmarshalJSON parses JSON encoded data to a ProductItem struct
func (c *ProductItem) UnmarshalJSON(data []byte) error {
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

	_, ok = objmap["name"]
	if ok {
		err = json.Unmarshal(objmap["name"], &c.Name)
		if err != nil {
			return err
		}
	}

	_, ok = objmap["product_id"]
	if ok {
		err = json.Unmarshal(objmap["product_id"], &c.ProductID)
		if err != nil {
			return err
		}
	}

	_, ok = objmap["variation_id"]
	if ok {
		err = json.Unmarshal(objmap["variation_id"], &c.VariationID)
		if err != nil {
			return err
		}
	}

	_, ok = objmap["quantity"]
	if ok {
		err = json.Unmarshal(objmap["quantity"], &c.Quantity)
		if err != nil {
			return err
		}
	}

	_, ok = objmap["sku"]
	if ok {
		err = json.Unmarshal(objmap["sku"], &c.SKU)
		if err != nil {
			return err
		}
	}

	_, ok = objmap["tax_class"]
	if ok {
		err = json.Unmarshal(objmap["tax_class"], &c.TaxClass)
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

	_, ok = objmap["subtotal"]
	if ok {
		err = json.Unmarshal(objmap["subtotal"], &c.Subtotal)
		if err != nil {
			c.Subtotal, err = tryMarshalStringAsFloat64(objmap["subtotal"])
			if err != nil {
				return err
			}
		}
	}

	_, ok = objmap["subtotal_tax"]
	if ok {
		err = json.Unmarshal(objmap["subtotal_tax"], &c.SubtotalTax)
		if err != nil {
			c.SubtotalTax, err = tryMarshalStringAsFloat64(objmap["subtotal_tax"])
			if err != nil {
				return err
			}
		}
	}

	_, ok = objmap["price"]
	if ok {
		err = json.Unmarshal(objmap["price"], &c.Price)
		if err != nil {
			c.Price, err = tryMarshalStringAsFloat64(objmap["price"])
			if err != nil {
				return err
			}
		}
	}

	// embedded *TaxableCost

	var tc TaxableCost
	errTC := json.Unmarshal(data, &tc)
	if errTC != nil {
		return errTC
	}
	c.TaxableCost = &tc

	return nil
}
