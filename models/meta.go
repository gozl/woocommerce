package models

// MetaItem represents an arbitary metadata entry
type MetaItem struct {
	ID int       `json:"id,omitempty"`
	Name string  `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}
