package yeelight

import (
	"context"
	"time"
)

type AdjustAction string

const (
	AdjustActionIncrease AdjustAction = "increase"
	AdjustActionDecrease AdjustAction = "decrease"
	AdjustActionCircle   AdjustAction = "circle"
)

type AdjustProp string

const (
	AdjustPropBright           AdjustProp = "bright"
	AdjustPropColorTemperature AdjustProp = "ct"
	AdjustPropColor            AdjustProp = "color"
)

func (c Client) SetAdjust(ctx context.Context, action AdjustAction, prop AdjustProp) error {
	return c.setAdjust(ctx, MethodSetAdjust, action, prop)
}

func (c Client) SetBackgroundAdjust(ctx context.Context, action AdjustAction, prop AdjustProp) error {
	return c.setAdjust(ctx, MethodSetBgAdjust, action, prop)
}

func (c Client) AdjustBright(ctx context.Context, percentage int, duration time.Duration) error {
	return c.adjustValue(ctx, MethodAdjustBright, percentage, duration)
}

func (c Client) AdjustBackgroundBright(ctx context.Context, percentage int, duration time.Duration) error {
	return c.adjustValue(ctx, MethodBgAdjustBright, percentage, duration)
}

func (c Client) AdjustColorTemperature(ctx context.Context, percentage int, duration time.Duration) error {
	return c.adjustValue(ctx, MethodAdjustCt, percentage, duration)
}

func (c Client) AdjustBackgroundColorTemperature(ctx context.Context, percentage int, duration time.Duration) error {
	return c.adjustValue(ctx, MethodBgAdjustCt, percentage, duration)
}

func (c Client) AdjustColor(ctx context.Context, percentage int, duration time.Duration) error {
	return c.adjustValue(ctx, MethodAdjustColor, percentage, duration)
}

func (c Client) AdjustBackgroundColor(ctx context.Context, percentage int, duration time.Duration) error {
	return c.adjustValue(ctx, MethodBgAdjustColor, percentage, duration)
}

func (c Client) setAdjust(ctx context.Context, method string, action AdjustAction, prop AdjustProp) error {
	return c.rawWithOk(ctx, method, action, prop)
}

func (c Client) adjustValue(ctx context.Context, method string, percentage int, duration time.Duration) error {
	if err := ValidatePercentage(percentage); err != nil {
		return err
	}

	if err := ValidateDuration(duration); err != nil {
		return err
	}

	return c.rawWithOk(ctx, method, percentage, duration.Milliseconds())
}
