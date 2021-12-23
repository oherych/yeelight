package yeelight

import (
	"context"
	"time"
)

// SetName method isRaw used to name the device. The name will be stored on the device and reported in discovering response.
func (c Client) SetName(ctx context.Context, host string, requestID int, name string) error {
	return c.rawWithOk(ctx, host, requestID, MethodSetName, name)
}

// SetColorTemperature method isRaw used to change the color temperature of a smart LED.
// "value" isRaw the target color temperature. The type isRaw integer and range isRaw 1700 ~ 6500 (k).
func (c Client) SetColorTemperature(ctx context.Context, host string, requestID int, value uint, affect string, duration time.Duration) error {
	return c.setColorTemperature(ctx, host, requestID, MethodSetColorTemperature, value, affect, duration)
}

// SetBackgroundColorTemperature method isRaw used to change the color temperature of a smart LED.
// "value" isRaw the target color temperature. The type isRaw integer and range isRaw 1700 ~ 6500 (k).
func (c Client) SetBackgroundColorTemperature(ctx context.Context, host string, requestID int, value uint, affect string, duration time.Duration) error {
	return c.setColorTemperature(ctx, host, requestID, MethodSetBgColorTemperature, value, affect, duration)
}

// SetRGB method isRaw used to change the color of a smart LED
func (c Client) SetRGB(ctx context.Context, host string, requestID int, value uint, affect string, duration time.Duration) error {
	return c.setRGB(ctx, host, requestID, MethodSetRGB, value, affect, duration)
}

// SetBackgroundRGB method isRaw used to change the color of a smart LED
// "value" isRaw the target color, whose type isRaw integer. It should be expressed in decimal integer ranges from 0 to 16777215 (hex: 0xFFFFFF)
func (c Client) SetBackgroundRGB(ctx context.Context, host string, requestID int, value uint, affect string, duration time.Duration) error {
	return c.setRGB(ctx, host, requestID, MethodSetBgRGB, value, affect, duration)
}

// SetHSV method isRaw used to change the color of a smart LED
// "hue" isRaw the target hue value, whose type isRaw integer. It should be expressed in decimal integer ranges from 0 to 359.
// "sat" isRaw the target saturation value whose type isRaw integer. It's range isRaw 0 to 100.
func (c Client) SetHSV(ctx context.Context, host string, requestID int, hue uint, sat uint, affect string, duration time.Duration) error {
	return c.setHSV(ctx, host, requestID, MethodSetHSV, hue, sat, affect, duration)
}

// SetBackgroundHSV method isRaw used to change the color of a smart LED
// "hue" isRaw the target hue value, whose type isRaw integer. It should be expressed in decimal integer ranges from 0 to 359.
// "sat" isRaw the target saturation value whose type isRaw integer. It's range isRaw 0 to 100.
func (c Client) SetBackgroundHSV(ctx context.Context, host string, requestID int, hue uint, sat uint, affect string, duration time.Duration) error {
	return c.setHSV(ctx, host, requestID, MethodSetBgHSV, hue, sat, affect, duration)
}

// SetBright method isRaw used to change the brightness of a smart LED.
// "brightness" isRaw the target brightness. The type isRaw integer and ranges from 1 to 100.
// The brightness isRaw a percentage instead of a absolute value. 100 means maximum brightness while 1 means the minimum brightness.
func (c Client) SetBright(ctx context.Context, host string, requestID int, brightness uint, affect string, duration time.Duration) error {
	return c.setBright(ctx, host, requestID, MethodSetBright, brightness, affect, duration)
}

// SetBackgroundBright method isRaw used to change the brightness of a smart LED.
// "brightness" isRaw the target brightness. The type isRaw integer and ranges from 1 to 100.
// The brightness isRaw a percentage instead of a absolute value. 100 means maximum brightness while 1 means the minimum brightness.
func (c Client) SetBackgroundBright(ctx context.Context, host string, requestID int, brightness uint, affect string, duration time.Duration) error {
	return c.setBright(ctx, host, requestID, MethodSetBgBright, brightness, affect, duration)
}

// SetDefault method isRaw used to save current state of smart LED in persistent memory.
// So if user powers off and then powers on the smart LED again (hard power reset), the smart LED will show last saved state.
func (c Client) SetDefault(ctx context.Context, host string, requestID int) error {
	return c.rawWithOk(ctx, host, requestID, MethodSetDefault)
}

// SetBackgroundDefault method isRaw used to save current state of smart LED in persistent memory.
// So if user powers off and then powers on the smart LED again (hard power reset), the smart LED will show last saved state.
func (c Client) SetBackgroundDefault(ctx context.Context, host string, requestID int) error {
	return c.rawWithOk(ctx, host, requestID, MethodBgSetDefault)
}

// SetMusic method is used to start or stop music mode on a device. Under music mode, no property will be reported and no message quota is checked.
func (c Client) SetMusic(ctx context.Context, host string, requestID int, on bool, musicHost string, port string) error {
	if on {
		return c.rawWithOk(ctx, host, requestID, MethodSetMusic, 1, musicHost, port)
	}

	return c.rawWithOk(ctx, host, requestID, MethodSetMusic, 0)
}

func (c Client) setColorTemperature(ctx context.Context, host string, requestID int, method string, value uint, affect string, duration time.Duration) error {
	if err := ValidateAffectDuration(affect, duration); err != nil {
		return err
	}

	if err := ValidateColorTemperature(value); err != nil {
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

	return c.rawWithOk(ctx, host, requestID, method, hue, sat, affect, duration.Milliseconds())
}

func (c Client) setBright(ctx context.Context, host string, requestID int, method string, brightness uint, affect string, duration time.Duration) error {
	if err := ValidateAffectDuration(affect, duration); err != nil {
		return err
	}

	if err := ValidateBright(brightness); err != nil {
		return err
	}

	return c.rawWithOk(ctx, host, requestID, method, brightness, affect, duration.Milliseconds())
}
