package yeelight

import (
	"context"
	"time"
)

func (c Client) Power(ctx context.Context, host string, requestID int, on bool, affect string, duration time.Duration) error {
	if err := ValidateAffectDuration(affect, duration); err != nil {
		return err
	}

	return c.rawWithOk(ctx, host, requestID, MethodSetPower, c.offOn(on), affect, duration.Milliseconds())
}

func (c Client) Toggle(ctx context.Context, host string, requestID int) error {
	return c.rawWithOk(ctx, host, requestID, MethodToggle)
}

func (c Client) offOn(on bool) string {
	if on {
		return "on"
	}

	return "off"
}
