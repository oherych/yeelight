package yeelight

import (
	"time"
)

func ValidateAffectDuration(affect string, duration time.Duration) error {
	if !IsAffect(affect) {
		return ErrWrongAffect
	}

	if err := ValidateDuration(duration); err != nil {
		return err
	}

	return nil
}

func ValidateDuration(duration time.Duration) error {
	if duration < (30 * time.Millisecond) {
		return ErrDurationTooSmall
	}

	return nil
}

func ValidatePercentage(value int) error {
	if value < 0 || value > 100 {
		return ErrWrongPercentage
	}

	return nil
}

func ValidateRGB(value int) error {
	const maxRGB = 16777215 // (hex: 0xFFFFFF)

	if value > maxRGB {
		return ErrWrongRGBValue
	}

	return nil
}

func ValidateHue(value int) error {
	const maxHue = 359

	if value > maxHue {
		return ErrWrongHueValue
	}

	return nil
}

func ValidateSat(value int) error {
	const maxSat = 100

	if value > maxSat {
		return ErrWrongSatValue
	}

	return nil
}

func ValidateBright(value int) error {
	if value < 0 || value > 100 {
		return ErrWrongBrightValue
	}

	return nil
}

func ValidateColorTemperature(value int) error {
	if value < 1700 || value > 6500 {
		return ErrWrongColorTemperature
	}

	return nil
}
