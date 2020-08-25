package models

import (
	"encoding/json"
	"strconv"
	"errors"
)

func tryMarshalStringAsFloat64(m json.RawMessage) (float64, error) {
	var value string
	err := json.Unmarshal(m, &value)
	if err != nil {
		return 0, err
	}

	retval, errParse := strconv.ParseFloat(value, 64)
	if errParse != nil {
		return 0, errParse
	}
	return retval, nil
}

func tryMarshalStringAsInt(m json.RawMessage) (int, error) {
	var value string
	err := json.Unmarshal(m, &value)
	if err != nil {
		return 0, err
	}

	retval, errParse := strconv.Atoi(value)
	if errParse != nil {
		return 0, errParse
	}
	return retval, nil
}

func tryMarshalMetadata(m json.RawMessage) ([]MetaItem, error) {
	var p []MetaItem
	err := json.Unmarshal(m, &p)
	if err == nil {
		return p, nil
	}

	var s string
	err = json.Unmarshal(m, &s)
	if err != nil {
		return nil, err
	}

	if len(s) == 0 {
		return []MetaItem{}, nil
	}

	return nil, errors.New("metaitem is non-empty string: " + s)
}