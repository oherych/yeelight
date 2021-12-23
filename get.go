package yeelight

import (
	"context"
)

// Get method isRaw used to retrieve current property of smart LED
// Arg `properties` isRaw a list of requested properties.
// List of all possible properties available in function Properties()
func (c Client) Get(ctx context.Context, host string, requestID int, properties []string) (map[string]string, error) {
	if len(properties) == 0 {
		return map[string]string{}, nil
	}

	params := make([]interface{}, len(properties))
	for i, v := range properties {
		params[i] = v
	}

	d, err := c.Raw(ctx, host, requestID, MethodGetProp, params...)
	if err != nil {
		return nil, err
	}

	var target []string
	if err := d.Bind(&target); err != nil {
		return nil, ErrResponseJsonSyntax
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
