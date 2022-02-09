package yeelight

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValidateDuration(t *testing.T) {
	tests := map[time.Duration]error{
		MinDuration - 1: ErrDurationTooSmall,
		MinDuration:     nil,
		MinDuration + 1: nil,
	}

	for duration, exp := range tests {
		t.Run(duration.String(), func(t *testing.T) {
			err := ValidateDuration(duration)

			assert.Equal(t, exp, err)
		})
	}
}

func TestValidateAdjustPercentage(t *testing.T) {
	tests := map[int]error{
		MinAdjustPercentage - 1: ErrWrongAdjustPercentage,
		MinAdjustPercentage:     nil,
		MinAdjustPercentage + 1: nil,
		MaxAdjustPercentage:     nil,
		MaxAdjustPercentage + 1: ErrWrongAdjustPercentage,
	}

	for percentage, exp := range tests {
		t.Run(strconv.Itoa(percentage), func(t *testing.T) {
			err := ValidateAdjustPercentage(percentage)

			assert.Equal(t, exp, err)
		})
	}
}

func TestValidateRGB(t *testing.T) {
	tests := map[int]error{
		MinRGB - 1: ErrWrongRGBValue,
		MinRGB:     nil,
		MaxRGB:     nil,
		MaxRGB + 1: ErrWrongRGBValue,
	}

	for rgb, exp := range tests {
		t.Run(strconv.Itoa(rgb), func(t *testing.T) {
			err := ValidateRGB(rgb)

			assert.Equal(t, exp, err)
		})
	}
}

func TestValidateHue(t *testing.T) {
	tests := map[int]error{
		MinHue - 1: ErrWrongHueValue,
		MinHue:     nil,
		MaxHue:     nil,
		MaxHue + 1: ErrWrongHueValue,
	}

	for hue, exp := range tests {
		t.Run(strconv.Itoa(hue), func(t *testing.T) {
			err := ValidateHue(hue)

			assert.Equal(t, exp, err)
		})
	}
}

func TestValidateSat(t *testing.T) {
	tests := map[int]error{
		MinSat - 1: ErrWrongSatValue,
		MinSat:     nil,
		MaxSat:     nil,
		MaxSat + 1: ErrWrongSatValue,
	}

	for hue, exp := range tests {
		t.Run(strconv.Itoa(hue), func(t *testing.T) {
			err := ValidateSat(hue)

			assert.Equal(t, exp, err)
		})
	}
}

func TestValidateBright(t *testing.T) {
	tests := map[int]error{
		MinBright - 1: ErrWrongBrightValue,
		MinBright:     nil,
		MaxBright:     nil,
		MaxBright + 1: ErrWrongBrightValue,
	}

	for hue, exp := range tests {
		t.Run(strconv.Itoa(hue), func(t *testing.T) {
			err := ValidateBright(hue)

			assert.Equal(t, exp, err)
		})
	}
}

func TestValidateColorTemperature(t *testing.T) {
	tests := map[int]error{
		MinColorTemperature - 1: ErrWrongColorTemperature,
		MinColorTemperature:     nil,
		MaxColorTemperature:     nil,
		MaxColorTemperature + 1: ErrWrongColorTemperature,
	}

	for hue, exp := range tests {
		t.Run(strconv.Itoa(hue), func(t *testing.T) {
			err := ValidateColorTemperature(hue)

			assert.Equal(t, exp, err)
		})
	}
}
