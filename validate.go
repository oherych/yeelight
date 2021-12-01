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
