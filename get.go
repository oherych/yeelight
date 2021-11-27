package yeelight

import (
	"context"
	"encoding/json"
	"errors"
)

var (
	// ErrWrongNumberOfResultItems says that response has wrong number of result items
	ErrWrongNumberOfResultItems = errors.New("wrong number of result items")
)

// Get method is used to retrieve current property of smart LED
// Arg `properties` is a list of requested properties.
// List of all possible properties available in function Properties()
func (c Client) Get(ctx context.Context, host string, requestID int, properties []string) (map[string]string, error) {
	params := make([]interface{}, len(properties))
	for i, v := range properties {
		params[i] = v
	}

	d, err := c.Raw(ctx, host, requestID, MethodGetProp, params...)
	if err != nil {
		return nil, err
	}

	var target []string
	if err := json.Unmarshal(d.Result, &target); err != nil {
		return nil, err
	}

	if len(properties) != len(target) {
		return nil, ErrWrongNumberOfResultItems
	}

	result := make(map[string]string, len(properties))
	for i, key := range properties {
		result[key] = target[i]
	}

	return result, nil
}
