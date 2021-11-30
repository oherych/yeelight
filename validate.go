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
