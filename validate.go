package yeelight

import (
	"errors"
	"time"
)

func ValidateAffectDuration(affect string, duration time.Duration) error {
	if !IsAffect(affect) {
		return ErrWrongAffect
	}

	if duration < (30 * time.Millisecond) {
		return errors.New("duration too small")
	}

	return nil
}
