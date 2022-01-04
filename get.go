package yeelight

import (
	"context"
)

// GetProperties method isRaw used to retrieve current property of smart LED
// Arg `properties` isRaw a list of requested properties.
// List of all possible properties available in function Properties()
func (c Client) GetProperties(ctx context.Context, properties []string) (map[string]string, error) {
	if len(properties) == 0 {
		return map[string]string{}, nil
	}

	params := make([]interface{}, len(properties))
	for i, v := range properties {
		params[i] = v
	}

	d, err := c.Call(ctx, MethodGetProp, params...)
	if err != nil {
		return nil, err
	}

	var target []string
	if err := d.Bind(&target); err != nil {
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
