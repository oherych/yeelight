package yeelight

import (
	"context"
	"fmt"
	"time"
)

const (
	FlowModeColor        = 1
	FlowColorTemperature = 2
	FlowSleep            = 7
)

type FlowExpression struct {
	Duration   time.Duration
	Mode       int
	Value      int
	Brightness int
}

// StartColorFlow method is used to start a color flow. Color flow is a series of smart LED visible state changing.
// It can be brightness changing, color changing or color temperature changing.
func (c Client) StartColorFlow(ctx context.Context, host string, requestID int, count uint, action uint, expressions []FlowExpression) error {
	return c.startColorFlow(ctx, host, requestID, MethodStartCF, count, action, expressions)
}

// StartBackgroundColorFlow method is used to start a color flow. Color flow is a series of smart LED visible state changing.
// It can be brightness changing, color changing or color temperature changing.
func (c Client) StartBackgroundColorFlow(ctx context.Context, host string, requestID int, count uint, action uint, expressions []FlowExpression) error {
	return c.startColorFlow(ctx, host, requestID, MethodBgStartCF, count, action, expressions)
}

func (c Client) StopColorFlow(ctx context.Context, host string, requestID int) error {
	return c.stopColorFlow(ctx, host, requestID, MethodStopCF)
}

func (c Client) StopBackgroundColorFlow(ctx context.Context, host string, requestID int) error {
	return c.stopColorFlow(ctx, host, requestID, MethodBgStopCF)
}

func (c Client) startColorFlow(ctx context.Context, host string, requestID int, method string, count uint, action uint, expressions []FlowExpression) error {
	var expressionStr string
	for _, exp := range expressions {
		if len(expressionStr) > 0 {
			expressionStr += ","
		}

		expressionStr += fmt.Sprintf("%d,%d,%d,%d", exp.Duration.Milliseconds(), exp.Mode, exp.Value, exp.Brightness)
	}

	return c.rawWithOk(ctx, host, requestID, method, count, action, expressionStr)
}

func (c Client) stopColorFlow(ctx context.Context, host string, requestID int, method string) error {
	return c.rawWithOk(ctx, host, requestID, method)
}
