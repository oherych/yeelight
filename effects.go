package yeelight

import "errors"

// Affects
const (
	AffectSudden = "sudden"
	AffectSmooth = "smooth"
)

var (
	// ErrWrongAffect wrong affect name
	ErrWrongAffect = errors.New("wrong affect name")

	ErrDurationTooSmall      = errors.New("duration too small")
	ErrWrongRGBValue         = errors.New("wrong RGB value")
	ErrWrongHueValue         = errors.New("wrong Hue value")
	ErrWrongSatValue         = errors.New("wrong Sat value")
	ErrWrongBrightValue      = errors.New("wrong Bright value")
	ErrWrongColorTemperature = errors.New("wrong color temperature")
)

// IsAffect return true if affect name is correct
func IsAffect(name string) bool {
	return name == AffectSudden || name == AffectSmooth
}

// Affects return list of all supported affects
func Affects() []string {
	return []string{AffectSudden, AffectSmooth}
}
