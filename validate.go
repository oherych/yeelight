package yeelight

import (
	"time"
)

func ValidateAffectDuration(affect string, duration time.Duration) error {
	if !IsAffect(affect) {
		return ErrWrongAffect
	}

	if duration < (30 * time.Millisecond) {
		return ErrDurationTooSmall
	}

	return nil
}

func ValidateRGB(value uint) error {
	const maxRGB = 16777215 // (hex: 0xFFFFFF)

	if value > maxRGB {
		return ErrWrongRGBValue
	}

	return nil
}

func ValidateHue(value uint) error {
	const maxHue = 359

	if value > maxHue {
		return ErrWrongHueValue
	}

	return nil
}

func ValidateSat(value uint) error {
	const maxSat = 100

	if value > maxSat {
		return ErrWrongSatValue
	}

	return nil
}

func ValidateBright(value uint) error {
	if value < 0 || value > 100 {
		return ErrWrongBrightValue
	}

	return nil
}

func ValidateColorTemperature(value uint) error {
	if value < 1700 || value > 6500 {
		return ErrWrongColorTemperature
	}

	return nil
}
