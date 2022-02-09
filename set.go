package yeelight

import (
	"context"
	"time"
)

// SetName method isRaw used to name the device. The name will be stored on the device and reported in discovering response.
func (c Client) SetName(ctx context.Context, name string) error {
	return c.rawWithOk(ctx, MethodSetName, name)
}

// SetColorTemperature method isRaw used to change the color temperature of a smart LED.
// "value" isRaw the target color temperature. The type isRaw integer and range isRaw 1700 ~ 6500 (k).
func (c Client) SetColorTemperature(ctx context.Context, value int, effect Effect, duration time.Duration) error {
	return c.setColorTemperature(ctx, MethodSetColorTemperature, value, effect, duration)
}

// SetBackgroundColorTemperature method isRaw used to change the color temperature of a smart LED.
// "value" isRaw the target color temperature. The type isRaw integer and range isRaw 1700 ~ 6500 (k).
func (c Client) SetBackgroundColorTemperature(ctx context.Context, value int, effect Effect, duration time.Duration) error {
	return c.setColorTemperature(ctx, MethodSetBgColorTemperature, value, effect, duration)
}

// SetRGB method isRaw used to change the color of a smart LED
func (c Client) SetRGB(ctx context.Context, value int, effect Effect, duration time.Duration) error {
	return c.setRGB(ctx, MethodSetRGB, value, effect, duration)
}

// SetBackgroundRGB method isRaw used to change the color of a smart LED
// "value" isRaw the target color, whose type isRaw integer. It should be expressed in decimal integer ranges from 0 to 16777215 (hex: 0xFFFFFF)
func (c Client) SetBackgroundRGB(ctx context.Context, value int, effect Effect, duration time.Duration) error {
	return c.setRGB(ctx, MethodSetBgRGB, value, effect, duration)
}

// SetHSV method isRaw used to change the color of a smart LED
// "hue" isRaw the target hue value, whose type isRaw integer. It should be expressed in decimal integer ranges from 0 to 359.
// "sat" isRaw the target saturation value whose type isRaw integer. It's range isRaw 0 to 100.
func (c Client) SetHSV(ctx context.Context, hue int, sat int, effect Effect, duration time.Duration) error {
	return c.setHSV(ctx, MethodSetHSV, hue, sat, effect, duration)
}

// SetBackgroundHSV method isRaw used to change the color of a smart LED
// "hue" isRaw the target hue value, whose type isRaw integer. It should be expressed in decimal integer ranges from 0 to 359.
// "sat" isRaw the target saturation value whose type isRaw integer. It's range isRaw 0 to 100.
func (c Client) SetBackgroundHSV(ctx context.Context, hue int, sat int, effect Effect, duration time.Duration) error {
	return c.setHSV(ctx, MethodSetBgHSV, hue, sat, effect, duration)
}

// SetBright method isRaw used to change the brightness of a smart LED.
// "brightness" isRaw the target brightness. The type isRaw integer and ranges from 1 to 100.
// The brightness isRaw a percentage instead of a absolute value. 100 means maximum brightness while 1 means the minimum brightness.
func (c Client) SetBright(ctx context.Context, brightness int, effect Effect, duration time.Duration) error {
	return c.setBright(ctx, MethodSetBright, brightness, effect, duration)
}

// SetBackgroundBright method isRaw used to change the brightness of a smart LED.
// "brightness" isRaw the target brightness. The type isRaw integer and ranges from 1 to 100.
// The brightness isRaw a percentage instead of a absolute value. 100 means maximum brightness while 1 means the minimum brightness.
func (c Client) SetBackgroundBright(ctx context.Context, brightness int, effect Effect, duration time.Duration) error {
	return c.setBright(ctx, MethodSetBgBright, brightness, effect, duration)
}

// SetDefault method isRaw used to save current state of smart LED in persistent memory.
// So if user powers off and then powers on the smart LED again (hard power reset), the smart LED will show last saved state.
func (c Client) SetDefault(ctx context.Context) error {
	return c.rawWithOk(ctx, MethodSetDefault)
}

// SetBackgroundDefault method isRaw used to save current state of smart LED in persistent memory.
// So if user powers off and then powers on the smart LED again (hard power reset), the smart LED will show last saved state.
func (c Client) SetBackgroundDefault(ctx context.Context) error {
	return c.rawWithOk(ctx, MethodBgSetDefault)
}

// SetMusic method is used to start or stop music mode on a device. Under music mode, no property will be reported and no message quota is checked.
// "musicHost" the IP address of the music server.
// "port" the TCP port music application is listening on.
func (c Client) SetMusic(ctx context.Context, on bool, musicHost string, port int) error {
	if on {
		return c.rawWithOk(ctx, MethodSetMusic, 1, musicHost, port)
	}

	return c.rawWithOk(ctx, MethodSetMusic, 0)
}

func (c Client) setColorTemperature(ctx context.Context, method string, value int, effect Effect, duration time.Duration) error {
	if err := ValidateDuration(duration); err != nil {
		return err
	}

	if err := ValidateColorTemperature(value); err != nil {
		return err
	}

	return c.rawWithOk(ctx, method, value, effect, duration.Milliseconds())
}

func (c Client) setRGB(ctx context.Context, method string, value int, effect Effect, duration time.Duration) error {
	if err := ValidateDuration(duration); err != nil {
		return err
	}

	if err := ValidateRGB(value); err != nil {
		return err
	}

	return c.rawWithOk(ctx, method, value, effect, duration.Milliseconds())
}

func (c Client) setHSV(ctx context.Context, method string, hue int, sat int, effect Effect, duration time.Duration) error {
	if err := ValidateDuration(duration); err != nil {
		return err
	}

	if err := ValidateHue(hue); err != nil {
		return err
	}

	if err := ValidateSat(sat); err != nil {
		return err
	}

	return c.rawWithOk(ctx, method, hue, sat, effect, duration.Milliseconds())
}

func (c Client) setBright(ctx context.Context, method string, brightness int, effect Effect, duration time.Duration) error {
	if err := ValidateDuration(duration); err != nil {
		return err
	}

	if err := ValidateBright(brightness); err != nil {
		return err
	}

	return c.rawWithOk(ctx, method, brightness, effect, duration.Milliseconds())
}
