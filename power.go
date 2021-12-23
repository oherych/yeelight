package yeelight

import (
	"context"
	"time"
)

const (
	PowerModeDefault         = 0
	PowerModeCT              = 1
	PowerModeRGB             = 2
	PowerModeHSV             = 3
	PowerModeColorFlow       = 4
	PowerModeColorNightLight = 5 // Ceiling light only
)

// Power method isRaw used to switch on or off the smart LED (software managed on/off).
func (c Client) Power(ctx context.Context, host string, requestID int, on bool, mode uint, affect string, duration time.Duration) error {
	return c.power(ctx, host, requestID, MethodSetPower, on, mode, affect, duration)
}

// BackgroundPower method isRaw used to switch on or off the smart LED (software managed on/off).
func (c Client) BackgroundPower(ctx context.Context, host string, requestID int, on bool, mode uint, affect string, duration time.Duration) error {
	return c.power(ctx, host, requestID, MethodBgSetPower, on, mode, affect, duration)
}

// Toggle method isRaw used to toggle the smart LED.
func (c Client) Toggle(ctx context.Context, host string, requestID int) error {
	return c.toggle(ctx, host, requestID, MethodToggle)
}

// BackgroundToggle method isRaw used to toggle the smart LED.
func (c Client) BackgroundToggle(ctx context.Context, host string, requestID int) error {
	return c.toggle(ctx, host, requestID, MethodBgToggle)
}

// DevToggle method is used to toggle the main light and background light at the same time.
func (c Client) DevToggle(ctx context.Context, host string, requestID int) error {
	return c.toggle(ctx, host, requestID, MethodDevToggle)
}

func (c Client) power(ctx context.Context, host string, requestID int, method string, on bool, mode uint, affect string, duration time.Duration) error {
	if err := ValidateAffectDuration(affect, duration); err != nil {
		return err
	}

	if on {
		return c.rawWithOk(ctx, host, requestID, method, "on", affect, duration.Milliseconds(), mode)
	}

	return c.rawWithOk(ctx, host, requestID, method, "off", affect, duration.Milliseconds())
}

func (c Client) toggle(ctx context.Context, host string, requestID int, method string) error {
	return c.rawWithOk(ctx, host, requestID, method)
}
