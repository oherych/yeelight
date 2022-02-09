package yeelight

import (
	"time"
)

const (
	MinDuration = 30 * time.Millisecond

	MinAdjustPercentage = -100
	MaxAdjustPercentage = 100

	MinRGB = 0
	MaxRGB = 16777215 // (hex: 0xFFFFFF)

	MinHue = 0
	MaxHue = 359

	MinSat = 0
	MaxSat = 100

	MinBright = 0
	MaxBright = 100

	MinColorTemperature = 1700
	MaxColorTemperature = 6500
)

// ValidateDuration validate duration
// validation rules in MinDuration
func ValidateDuration(duration time.Duration) error {
	if duration < MinDuration {
		return ErrDurationTooSmall
	}

	return nil
}

// ValidateAdjustPercentage validate percentage for adjust_* methods
// validation rules in MinAdjustPercentage and MaxAdjustPercentage
func ValidateAdjustPercentage(value int) error {
	if value < MinAdjustPercentage || value > MaxAdjustPercentage {
		return ErrWrongAdjustPercentage
	}

	return nil
}

// ValidateRGB validate rgb
// validation rules in MinRGB and MaxRGB
func ValidateRGB(value int) error {
	if value < MinRGB || value > MaxRGB {
		return ErrWrongRGBValue
	}

	return nil
}

// ValidateHue validate hue
// validation rules in MinHue and MaxHue
func ValidateHue(value int) error {
	if value < MinHue || value > MaxHue {
		return ErrWrongHueValue
	}

	return nil
}

// ValidateSat validate sat
// validation rules in MinSat and MaxSat
func ValidateSat(value int) error {
	if value < MinSat || value > MaxSat {
		return ErrWrongSatValue
	}

	return nil
}

func ValidateBright(value int) error {
	if value < MinBright || value > MaxBright {
		return ErrWrongBrightValue
	}

	return nil
}

func ValidateColorTemperature(value int) error {
	if value < MinColorTemperature || value > MaxColorTemperature {
		return ErrWrongColorTemperature
	}

	return nil
}
