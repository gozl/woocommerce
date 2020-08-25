package models

import (
	"encoding/json"
)

// RefundItem represents a refund associated with an order
type RefundItem struct {
	ID int                 `json:"id,omitempty"`
	Reason string          `json:"reason,omitempty"`
	Total float64          `json:"total,omitempty"`
}

// UnmarshalJSON parses JSON encoded data to a RefundItem struct
func (c *RefundItem) UnmarshalJSON(data []byte) error {
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

	_, ok = objmap["reason"]
	if ok {
		err = json.Unmarshal(objmap["reason"], &c.Reason)
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

	return nil
}
