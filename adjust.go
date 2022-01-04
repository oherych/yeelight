package yeelight

import (
	"context"
	"time"
)

// SetAdjust method is used to change brightness, CT or color of a smart LED without knowing the current value, it's main used by controllers.
func (c Client) SetAdjust(ctx context.Context, action AdjustAction, prop AdjustProp) error {
	return c.setAdjust(ctx, MethodSetAdjust, action, prop)
}

// SetBackgroundAdjust method is used to change brightness, CT or color of a smart LED without knowing the current value, it's main used by controllers.
func (c Client) SetBackgroundAdjust(ctx context.Context, action AdjustAction, prop AdjustProp) error {
	return c.setAdjust(ctx, MethodSetBgAdjust, action, prop)
}

// AdjustBright method is used to adjust the brightness by specified percentage within specified duration.
func (c Client) AdjustBright(ctx context.Context, percentage int, duration time.Duration) error {
	return c.adjustValue(ctx, MethodAdjustBright, percentage, duration)
}

// AdjustBackgroundBright method is used to adjust the brightness by specified percentage within specified duration.
func (c Client) AdjustBackgroundBright(ctx context.Context, percentage int, duration time.Duration) error {
	return c.adjustValue(ctx, MethodBgAdjustBright, percentage, duration)
}

// AdjustColorTemperature method is used to adjust the color temperature by specified percentage within specified duration.
func (c Client) AdjustColorTemperature(ctx context.Context, percentage int, duration time.Duration) error {
	return c.adjustValue(ctx, MethodAdjustCt, percentage, duration)
}

// AdjustBackgroundColorTemperature method is used to adjust the color temperature by specified percentage within specified duration.
func (c Client) AdjustBackgroundColorTemperature(ctx context.Context, percentage int, duration time.Duration) error {
	return c.adjustValue(ctx, MethodBgAdjustCt, percentage, duration)
}

// AdjustColor method is used to adjust the color within specified duration.
func (c Client) AdjustColor(ctx context.Context, percentage int, duration time.Duration) error {
	return c.adjustValue(ctx, MethodAdjustColor, percentage, duration)
}

// AdjustBackgroundColor method is used to adjust the color within specified duration.
func (c Client) AdjustBackgroundColor(ctx context.Context, percentage int, duration time.Duration) error {
	return c.adjustValue(ctx, MethodBgAdjustColor, percentage, duration)
}

func (c Client) setAdjust(ctx context.Context, method string, action AdjustAction, prop AdjustProp) error {
	return c.rawWithOk(ctx, method, action, prop)
}

func (c Client) adjustValue(ctx context.Context, method string, percentage int, duration time.Duration) error {
	if err := validatePercentage(percentage); err != nil {
		return err
	}

	if err := validateDuration(duration); err != nil {
		return err
	}

	return c.rawWithOk(ctx, method, percentage, duration.Milliseconds())
}
