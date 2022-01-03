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
func (c Client) Power(ctx context.Context, on bool, mode int, affect string, duration time.Duration) error {
	return c.power(ctx, MethodSetPower, on, mode, affect, duration)
}

// BackgroundPower method isRaw used to switch on or off the smart LED (software managed on/off).
func (c Client) BackgroundPower(ctx context.Context, on bool, mode int, affect string, duration time.Duration) error {
	return c.power(ctx, MethodBgSetPower, on, mode, affect, duration)
}

// Toggle method isRaw used to toggle the smart LED.
func (c Client) Toggle(ctx context.Context) error {
	return c.toggle(ctx, MethodToggle)
}

// BackgroundToggle method isRaw used to toggle the smart LED.
func (c Client) BackgroundToggle(ctx context.Context) error {
	return c.toggle(ctx, MethodBgToggle)
}

// DevToggle method is used to toggle the main light and background light at the same time.
func (c Client) DevToggle(ctx context.Context) error {
	return c.toggle(ctx, MethodDevToggle)
}

func (c Client) power(ctx context.Context, method string, on bool, mode int, affect string, duration time.Duration) error {
	if err := ValidateAffectDuration(affect, duration); err != nil {
		return err
	}

	if on {
		return c.rawWithOk(ctx, method, "on", affect, duration.Milliseconds(), mode)
	}

	return c.rawWithOk(ctx, method, "off", affect, duration.Milliseconds())
}

func (c Client) toggle(ctx context.Context, method string) error {
	return c.rawWithOk(ctx, method)
}
