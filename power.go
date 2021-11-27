package yeelight

import (
	"context"
	"errors"
	"time"
)

func (c Client) Power(ctx context.Context, host string, requestID int, on bool, affect string, duration time.Duration) error {
	if !IsAffect(affect) {
		return ErrWrongAffect
	}

	var action string
	if on {
		action = "on"
	} else {
		action = "off"
	}

	d, err := c.Raw(ctx, host, requestID, MethodSetPower, action, affect, duration.Milliseconds())
	if err != nil {
		return err
	}

	if !d.IsOk() {
		return errors.New(d.String())
	}

	return nil
}

func (c Client) Toggle(ctx context.Context, host string, requestID int) error {
	d, err := c.Raw(ctx, host, requestID, MethodToggle)
	if err != nil {
		return err
	}

	if !d.IsOk() {
		return errors.New(d.String())
	}

	return nil
}
