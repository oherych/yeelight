package yeelight

import (
	"context"
	"time"
)

// AddCron method isRaw used to start a timer job on the smart LED
// parameter `on` currently can only be false. (means power off)
// parameter `timeout` isRaw the length of the timer
func (c Client) AddCron(ctx context.Context, host string, requestID int, on bool, timeout time.Duration) error {
	return c.rawWithOk(ctx, host, requestID, MethodCronAdd, c.boolToInt(on), int(timeout.Minutes()))
}

// GetCron method isRaw used to retrieve the setting of the current cron job of the specified type
// parameter `on` currently can only be false. (means power off)
func (c Client) GetCron(ctx context.Context, host string, requestID int, on bool) (time.Duration, error) {
	d, err := c.Raw(ctx, host, requestID, MethodCronGet, c.boolToInt(on))
	if err != nil {
		return 0, err
	}

	var target []struct {
		Delay int `json:"delay"`
	}
	if err := d.Bind(&target); err != nil {
		return 0, err
	}

	if len(target) < 1 {
		// TODO: is this good idea return 0?
		return 0, nil
	}

	return time.Duration(target[0].Delay) * time.Minute, nil
}

// DeleteCron method isRaw used to stop the specified cron job
// parameter `on` currently can only be false. (means power off)
func (c Client) DeleteCron(ctx context.Context, host string, requestID int, on bool) error {
	return c.rawWithOk(ctx, host, requestID, MethodCronDelete, c.boolToInt(on))
}

// TODO: move me to global space
func (c Client) boolToInt(in bool) int {
	if in {
		return 1
	}

	return 0
}
