package yeelight

import "errors"

var (
	// ErrWrongNumberOfResultItems says that response has wrong number of result items
	ErrWrongNumberOfResultItems = errors.New("wrong number of result items")
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

var (
	ErrMissingPortInAddress = errors.New("missing port in address")
	ErrConnect              = errors.New("connect error")
	ErrMethodNotSupported   = errors.New("method not supported")
	ErrResponseJsonSyntax   = errors.New("response json syntax error")
)

type UnknownError string

func (e UnknownError) Error() string {
	return string(e)
}

