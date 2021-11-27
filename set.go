package yeelight

import (
	"context"
	"errors"
	"time"
)

func (c Client) SetName(ctx context.Context, host string, requestID int, name string) error {
	d, err := c.Raw(ctx, host, requestID, MethodSetName, name)
	if err != nil {
		return err
	}

	if !d.IsOk() {
		return errors.New(d.String())
	}

	return nil
}

// SetColorTemperature method is used to change the color temperature of a smart LED
func (c Client) SetColorTemperature(ctx context.Context, host string, requestID int, value uint, affect string, duration time.Duration) error {
	return c.setColorTemperature(ctx, host, requestID, MethodSetColorTemperature, value, affect, duration)
}

// SetBackgroundColorTemperature method is used to change the color temperature of a smart LED
func (c Client) SetBackgroundColorTemperature(ctx context.Context, host string, requestID int, value uint, affect string, duration time.Duration) error {
	return c.setColorTemperature(ctx, host, requestID, MethodSetBgColorTemperature, value, affect, duration)
}

func (c Client) SetRGB(ctx context.Context, host string, requestID int, value uint, affect string, duration time.Duration) error {
	return c.setRGB(ctx, host, requestID, MethodSetRGB, value, affect, duration)
}

func (c Client) SetBackgroundRGB(ctx context.Context, host string, requestID int, value uint, affect string, duration time.Duration) error {
	return c.setRGB(ctx, host, requestID, MethodSetBgRGB, value, affect, duration)
}

func (c Client) SetHSV(ctx context.Context, host string, requestID int, hue uint, sat uint, affect string, duration time.Duration) error {
	return c.setHSV(ctx, host, requestID, MethodSetHSV, hue, sat, affect, duration)
}

func (c Client) SetBackgroundHSV(ctx context.Context, host string, requestID int, hue uint, sat uint, affect string, duration time.Duration) error {
	return c.setHSV(ctx, host, requestID, MethodSetBgHSV, hue, sat, affect, duration)
}

func (c Client) SetBright(ctx context.Context, host string, requestID int, value uint, affect string, duration time.Duration) error {
	return c.setBright(ctx, host, requestID, MethodSetBright, value, affect, duration)
}

func (c Client) SetBackgroundBright(ctx context.Context, host string, requestID int, value uint, affect string, duration time.Duration) error {
	return c.setBright(ctx, host, requestID, MethodSetBgBright, value, affect, duration)
}

func (c Client) SetDefault(ctx context.Context, host string, requestID int) error {
	d, err := c.Raw(ctx, host, requestID, MethodSetDefault)
	if err != nil {
		return err
	}

	return d.ToError()
}

func (c Client) setColorTemperature(ctx context.Context, host string, requestID int, method string, value uint, affect string, duration time.Duration) error {
	if err := ValidateAffectDuration(affect, duration); err != nil {
		return err
	}

	d, err := c.Raw(ctx, host, requestID, method, value, affect, duration.Milliseconds())
	if err != nil {
		return err
	}

	return d.ToError()
}

func (c Client) setRGB(ctx context.Context, host string, requestID int, method string, value uint, affect string, duration time.Duration) error {
	if !IsAffect(affect) {
		return ErrWrongAffect
	}

	d, err := c.Raw(ctx, host, requestID, method, value, affect, duration.Milliseconds())
	if err != nil {
		return err
	}

	if !d.IsOk() {
		return errors.New(d.String())
	}

	return nil
}

func (c Client) setHSV(ctx context.Context, host string, requestID int, method string, hue uint, sat uint, affect string, duration time.Duration) error {
	if !IsAffect(affect) {
		return ErrWrongAffect
	}

	d, err := c.Raw(ctx, host, requestID, method, hue, sat, affect, duration.Milliseconds())
	if err != nil {
		return err
	}

	if !d.IsOk() {
		return errors.New(d.String())
	}

	return nil
}

func (c Client) setBright(ctx context.Context, host string, requestID int, method string, value uint, affect string, duration time.Duration) error {
	if !IsAffect(affect) {
		return ErrWrongAffect
	}

	d, err := c.Raw(ctx, host, requestID, method, value, affect, duration.Milliseconds())
	if err != nil {
		return err
	}

	if !d.IsOk() {
		return errors.New(d.String())
	}

	return nil
}
