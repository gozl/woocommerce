package models

import (
	//"time"
	//"strconv"
	"encoding/json"
)

// Order reprents a product order
type Order struct {
	*PaymentTransaction
	ID int                    `json:"id,omitempty"`
	ParentID int              `json:"parent_id,omitempty"`
	Key string                `json:"order_key,omitempty"`
	CreatedBy string          `json:"created_via,omitempty"`
	Version string            `json:"version,omitempty"`
	Status string             `json:"status,omitempty"`
	CartHash string           `json:"cart_hash,omitempty"`
	Shipping ShippingInfo     `json:"shipping,omitempty"`
	Billing BillingInfo       `json:"billing,omitempty"`
	Created Timestamp         `json:"date_created_gmt,omitempty"`
	Updated Timestamp         `json:"date_modified_gmt,omitempty"`
	Completed Timestamp       `json:"date_completed_gmt,omitempty"`
	Coupons []CouponItem      `json:"coupon_lines,omitempty"`
	Refunds []RefundItem      `json:"refunds,omitempty"`
	Fees []FeeItem            `json:"fee_lines,omitempty"`
	Shipment []ShippingMethod `json:"shipping_lines,omitempty"`
	Taxes []TaxItem           `json:"tax_lines,omitempty"`
	Products []ProductItem    `json:"line_items,omitempty"`
	// Special
	Number int                `json:"number,omitempty"`
	Metadata []MetaItem       `json:"meta_data,omitempty"`
	// Ordering customer
	CustomerID int            `json:"customer_id,omitempty"`
	CustomerIP string         `json:"customer_ip_address,omitempty"`
	CustomerUA string         `json:"customer_user_agent,omitempty"`
	CustomerNote string       `json:"customer_note,omitempty"`
	// Price
	Currency string           `json:"currency,omitempty"`
	CurrencySymbol string     `json:"currency_symbol,omitempty"`
	Discount float64          `json:"discount_total,omitempty"`
	DiscountTax float64       `json:"discount_tax,omitempty"`
	ShippingCost float64      `json:"shipping_total,omitempty"`
	ShippingTax float64       `json:"shipping_tax,omitempty"`
	CartTax float64           `json:"cart_tax,omitempty"`
	Total float64             `json:"total,omitempty"`
	TotalTax float64          `json:"total_tax,omitempty"`
	TaxIncluded bool          `json:"prices_include_tax,omitempty"`
}


// UnmarshalJSON parses JSON encoded data to an Order struct
func (c *Order) UnmarshalJSON(data []byte) error {
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

	_, ok = objmap["number"]
	if ok {
		err = json.Unmarshal(objmap["number"], &c.Number)
		if err != nil {
			c.Number, err = tryMarshalStringAsInt(objmap["number"])
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

	// embedded *PaymentTransaction

	var pt PaymentTransaction
	errPT := json.Unmarshal(data, &pt)
	if errPT != nil {
		return errPT
	}
	c.PaymentTransaction = &pt

	// as-is fields

/*
$t = @'
    _, ok = objmap["<<0>>"]
    if ok {
        err = json.Unmarshal(objmap["<<0>>"], &c.<<1>>)
        if err != nil {
            return err
        }
	}
	
'@
$a | % {
	$x = $_ -split 'json' | select -First 1
	$x = $x.trim() -split ' ' | select -first 1
	$y = $_ -split 'json' | select -Last 1
	$y = $y -split ',' | select -first 1
	$y = $y -split '"' | select -last 1

	$t.Replace('<<0>>', $y).Replace('<<1>>', $x)
}
*/

	_, ok = objmap["customer_id"]
	if ok {
		err = json.Unmarshal(objmap["customer_id"], &c.CustomerID)
		if err != nil {
			return err
		}
	}

	_, ok = objmap["customer_ip_address"]
	if ok {
		err = json.Unmarshal(objmap["customer_ip_address"], &c.CustomerIP)
		if err != nil {
			return err
		}
	}

	_, ok = objmap["customer_user_agent"]
	if ok {
		err = json.Unmarshal(objmap["customer_user_agent"], &c.CustomerUA)
		if err != nil {
			return err
		}
	}

	_, ok = objmap["customer_note"]
	if ok {
		err = json.Unmarshal(objmap["customer_note"], &c.CustomerNote)
		if err != nil {
			return err
		}
	}

	    _, ok = objmap["id"]
    if ok {
        err = json.Unmarshal(objmap["id"], &c.ID)
        if err != nil {
            return err
        }
    }

    _, ok = objmap["parent_id"]
    if ok {
        err = json.Unmarshal(objmap["parent_id"], &c.ParentID)
        if err != nil {
            return err
        }
    }

    _, ok = objmap["order_key"]
    if ok {
        err = json.Unmarshal(objmap["order_key"], &c.Key)
        if err != nil {
            return err
        }
    }

    _, ok = objmap["created_via"]
    if ok {
        err = json.Unmarshal(objmap["created_via"], &c.CreatedBy)
        if err != nil {
            return err
        }
    }

    _, ok = objmap["version"]
    if ok {
        err = json.Unmarshal(objmap["version"], &c.Version)
        if err != nil {
            return err
        }
    }

    _, ok = objmap["status"]
    if ok {
        err = json.Unmarshal(objmap["status"], &c.Status)
        if err != nil {
            return err
        }
    }

    _, ok = objmap["cart_hash"]
    if ok {
        err = json.Unmarshal(objmap["cart_hash"], &c.CartHash)
        if err != nil {
            return err
        }
    }

    _, ok = objmap["shipping"]
    if ok {
        err = json.Unmarshal(objmap["shipping"], &c.Shipping)
        if err != nil {
            return err
        }
    }

    _, ok = objmap["billing"]
    if ok {
        err = json.Unmarshal(objmap["billing"], &c.Billing)
        if err != nil {
            return err
        }
    }

    _, ok = objmap["date_created_gmt"]
    if ok {
        err = json.Unmarshal(objmap["date_created_gmt"], &c.Created)
        if err != nil {
            return err
        }
    }

    _, ok = objmap["date_modified_gmt"]
    if ok {
        err = json.Unmarshal(objmap["date_modified_gmt"], &c.Updated)
        if err != nil {
            return err
        }
    }

    _, ok = objmap["date_completed_gmt"]
    if ok {
        err = json.Unmarshal(objmap["date_completed_gmt"], &c.Completed)
        if err != nil {
            return err
        }
    }

    _, ok = objmap["coupon_lines"]
    if ok {
        err = json.Unmarshal(objmap["coupon_lines"], &c.Coupons)
        if err != nil {
            return err
        }
    }

    _, ok = objmap["refunds"]
    if ok {
        err = json.Unmarshal(objmap["refunds"], &c.Refunds)
        if err != nil {
            return err
        }
    }

    _, ok = objmap["fee_lines"]
    if ok {
        err = json.Unmarshal(objmap["fee_lines"], &c.Fees)
        if err != nil {
            return err
        }
    }

    _, ok = objmap["shipping_lines"]
    if ok {
        err = json.Unmarshal(objmap["shipping_lines"], &c.Shipment)
        if err != nil {
            return err
        }
    }

    _, ok = objmap["tax_lines"]
    if ok {
        err = json.Unmarshal(objmap["tax_lines"], &c.Taxes)
        if err != nil {
            return err
        }
    }

    _, ok = objmap["line_items"]
    if ok {
        err = json.Unmarshal(objmap["line_items"], &c.Products)
        if err != nil {
            return err
        }
    }

	return nil
}
