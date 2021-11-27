package yeelight

import (
	"context"
	"time"
)

// AddCron method is used to start a timer job on the smart LED
// parameter `on` currently can only be false. (means power off)
// parameter `timeout` is the length of the timer
func (c Client) AddCron(ctx context.Context, host string, requestID int, on bool, timeout time.Duration) error {
	panic("implement me")
}

// GetCron method is used to retrieve the setting of the current cron job of the specified type
// parameter `on` currently can only be false. (means power off)
func (c Client) GetCron(ctx context.Context, host string, requestID int, on bool) (time.Duration, error) {
	panic("implement me")
}

// DeleteCron method is used to stop the specified cron job
// parameter `on` currently can only be false. (means power off)
func (c Client) DeleteCron(ctx context.Context, host string, requestID int, on bool) error {
	panic("implement me")
}
