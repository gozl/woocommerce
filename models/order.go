package models

import (
	//"time"
	"strconv"
	"encoding/json"
)

// OrderingCustomer represents basic customer info relating to an order
type OrderingCustomer struct {
	CustomerID int      `json:"customer_id,omitempty"`
	CustomerIP string   `json:"customer_ip_address,omitempty"`
	CustomerUA string   `json:"customer_user_agent,omitempty"`
	CustomerNote string `json:"customer_note,omitempty"`
}

// Order reprents a product order
type Order struct {
	//*Price
	*OrderingCustomer
	*PaymentTransaction
	ID int                    `json:"id,omitempty"`
	ParentID int              `json:"parent_id,omitempty"`
	Number int                `json:"number,omitempty"`
	Key string                `json:"order_key,omitempty"`
	CreatedBy string          `json:"created_via,omitempty"`
	Version string            `json:"version,omitempty"`
	Status string             `json:"status,omitempty"`
	CartHash string           `json:"cart_hash,omitempty"`
	Metadata []MetaItem       `json:"meta_data,omitempty"`
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
}

// MarshalJSON serializes a Order struct to JSON encoded data
func (c *Order) MarshalJSON() ([]byte, error) {
	type jsonObj Order
	return json.Marshal(&struct{
		Number        string `json:"number,omitempty"`
		*jsonObj
	}{
		Number: strconv.Itoa(c.Number),
		jsonObj:  (*jsonObj)(c),
	})
}

// UnmarshalJSON parses JSON encoded data to a Order struct
func (c *Order) UnmarshalJSON(data []byte) error {
	if data == nil || len(data) == 0 {
		return nil
	}

	type rObj Order
	aux := &struct{
		Number        string `json:"number,omitempty"`
		*rObj
	}{
		rObj: (*rObj)(c),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	var errInt error

	c.Number, errInt = strconv.Atoi(aux.Number)
	if errInt != nil {
		return errInt
	}

	return nil
}
