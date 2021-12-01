package yeelight

import (
	"context"
	"errors"
	"time"
)

// SetName method is used to name the device. The name will be stored on the device and reported in discovering response.
func (c Client) SetName(ctx context.Context, host string, requestID int, name string) error {
	return c.rawWithOk(ctx, host, requestID, MethodSetName, name)
}

// SetColorTemperature method is used to change the color temperature of a smart LED
func (c Client) SetColorTemperature(ctx context.Context, host string, requestID int, value uint, affect string, duration time.Duration) error {
	return c.setColorTemperature(ctx, host, requestID, MethodSetColorTemperature, value, affect, duration)
}

// SetBackgroundColorTemperature method is used to change the color temperature of a smart LED
func (c Client) SetBackgroundColorTemperature(ctx context.Context, host string, requestID int, value uint, affect string, duration time.Duration) error {
	return c.setColorTemperature(ctx, host, requestID, MethodSetBgColorTemperature, value, affect, duration)
}

// SetRGB method is used to change the color of a smart LED
func (c Client) SetRGB(ctx context.Context, host string, requestID int, value uint, affect string, duration time.Duration) error {
	return c.setRGB(ctx, host, requestID, MethodSetRGB, value, affect, duration)
}

// SetBackgroundRGB method is used to change the color of a smart LED
// "value" is the target color, whose type is integer. It should be expressed in decimal integer ranges from 0 to 16777215 (hex: 0xFFFFFF)
func (c Client) SetBackgroundRGB(ctx context.Context, host string, requestID int, value uint, affect string, duration time.Duration) error {
	return c.setRGB(ctx, host, requestID, MethodSetBgRGB, value, affect, duration)
}

// SetHSV method is used to change the color of a smart LED
// "hue" is the target hue value, whose type is integer. It should be expressed in decimal integer ranges from 0 to 359.
// "sat" is the target saturation value whose type is integer. It's range is 0 to 100.
func (c Client) SetHSV(ctx context.Context, host string, requestID int, hue uint, sat uint, affect string, duration time.Duration) error {
	return c.setHSV(ctx, host, requestID, MethodSetHSV, hue, sat, affect, duration)
}

// SetBackgroundHSV method is used to change the color of a smart LED
// "hue" is the target hue value, whose type is integer. It should be expressed in decimal integer ranges from 0 to 359.
// "sat" is the target saturation value whose type is integer. It's range is 0 to 100.
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
	return c.rawWithOk(ctx, host, requestID, MethodSetDefault)
}

func (c Client) setColorTemperature(ctx context.Context, host string, requestID int, method string, value uint, affect string, duration time.Duration) error {
	if err := ValidateAffectDuration(affect, duration); err != nil {
		return err
	}

	return c.rawWithOk(ctx, host, requestID, method, value, affect, duration.Milliseconds())
}

func (c Client) setRGB(ctx context.Context, host string, requestID int, method string, value uint, affect string, duration time.Duration) error {
	if err := ValidateAffectDuration(affect, duration); err != nil {
		return err
	}

	if err := ValidateRGB(value); err != nil {
		return err
	}

	return c.rawWithOk(ctx, host, requestID, method, value, affect, duration.Milliseconds())
}

func (c Client) setHSV(ctx context.Context, host string, requestID int, method string, hue uint, sat uint, affect string, duration time.Duration) error {
	if err := ValidateAffectDuration(affect, duration); err != nil {
		return err
	}

	if err := ValidateHue(hue); err != nil {
		return err
	}

	if err := ValidateSat(sat); err != nil {
		return err
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
	if err := ValidateAffectDuration(affect, duration); err != nil {
		return err
	}

	return c.rawWithOk(ctx, host, requestID, method, value, affect, duration.Milliseconds())
}
