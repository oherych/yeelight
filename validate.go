package yeelight

import (
	"time"
)

func validateAffectDuration(affect string, duration time.Duration) error {
	if !IsAffect(affect) {
		return ErrWrongAffect
	}

	if err := validateDuration(duration); err != nil {
		return err
	}

	return nil
}

func validateDuration(duration time.Duration) error {
	if duration < (30 * time.Millisecond) {
		return ErrDurationTooSmall
	}

	return nil
}

func validatePercentage(value int) error {
	if value < -100 || value > 100 {
		return ErrWrongPercentage
	}

	return nil
}

func validateRGB(value int) error {
	const maxRGB = 16777215 // (hex: 0xFFFFFF)

	if value < 0 || value > maxRGB {
		return ErrWrongRGBValue
	}

	return nil
}

func validateHue(value int) error {
	const maxHue = 359

	if value < 0 || value > maxHue {
		return ErrWrongHueValue
	}

	return nil
}

func validateSat(value int) error {
	const maxSat = 100

	if value < 0 || value > maxSat {
		return ErrWrongSatValue
	}

	return nil
}

func validateBright(value int) error {
	if value < 0 || value > 100 {
		return ErrWrongBrightValue
	}

	return nil
}

func validateColorTemperature(value int) error {
	if value < 1700 || value > 6500 {
		return ErrWrongColorTemperature
	}

	return nil
}
