package yeelight

import "errors"

var (
	// ErrWrongNumberOfResultItems says that response has wrong number of result items
	ErrWrongNumberOfResultItems = errors.New("wrong number of result items")

	// ErrCronIsUnset says that cron is unset
	ErrCronIsUnset = errors.New("cron is unset")

	// ErrDurationTooSmall says that duration too small
	ErrDurationTooSmall = errors.New("duration too small")

	// ErrWrongPercentage says that wrong percentage
	ErrWrongPercentage = errors.New("wrong percentage")

	// ErrWrongRGBValue says that wrong RGB value
	ErrWrongRGBValue = errors.New("wrong RGB value")

	// ErrWrongHueValue says that wrong Hue value
	ErrWrongHueValue = errors.New("wrong Hue value")

	// ErrWrongSatValue says that wrong Sat value
	ErrWrongSatValue = errors.New("wrong Sat value")

	// ErrWrongBrightValue says that wrong Bright value
	ErrWrongBrightValue = errors.New("wrong Bright value")

	// ErrWrongColorTemperature says that wrong color temperature
	ErrWrongColorTemperature = errors.New("wrong color temperature")

	// ErrRequiredMinimumOneExpression says that required minimum one expression
	ErrRequiredMinimumOneExpression = errors.New("required minimum one expression")

	// ErrMissingPortInAddress says that missing port in address
	ErrMissingPortInAddress = errors.New("missing port in address")

	// ErrConnect says that connect error
	ErrConnect = errors.New("connect error")

	// ErrMethodNotSupported says that method not supported
	ErrMethodNotSupported = errors.New("method not supported")

	// ErrResponseJSONSyntax says that response json has syntax error
	ErrResponseJSONSyntax = errors.New("response json has syntax error")
)

// UnknownError contain raw unknown error
type UnknownError string

func (e UnknownError) Error() string {
	return string(e)
}
