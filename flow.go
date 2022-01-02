package yeelight

import (
	"context"
	"fmt"
	"time"
)

// Color flow modes
const (
	FlowModeColor        = 1
	FlowColorTemperature = 2
	FlowSleep            = 7
)

// Color flow action
const (
	// FlowActionRecover means smart LED recover to the state before the color flow started.
	FlowActionRecover = 0

	// FlowActionStay means smart LED stay at the state when the flow is stopped.
	FlowActionStay = 1

	// FlowActionTurnOff means turn off the smart LED after the flow is stopped.
	FlowActionTurnOff = 2
)

// FlowExpression is the expression of the state changing series.
type FlowExpression struct {
	// Duration Gradual change time or sleep time. Minimum value 50 milliseconds.
	Duration time.Duration

	// Mode can be FlowModeColor, FlowColorTemperature, FlowSleep
	Mode int

	// Value is RGB value when mode is FlowModeColor, CT value when mode is FlowColorTemperature, Ignored when mode is FlowSleep.
	Value int

	// Brightness value, -1 or 1 ~ 100. Ignored when mode is FlowSleep.
	// When this value is -1, brightness in this tuple is ignored (only color or CT change takes effect).
	Brightness int
}

// StartColorFlow method is used to start a color flow. Color flow is a series of smart LED visible state changing.
// It can be brightness changing, color changing or color temperature changing.
// `count` is the total number of visible state changing before color flow stopped. 0 means infinite loop on the state changing.
func (c Client) StartColorFlow(ctx context.Context, count int, action int, expressions []FlowExpression) error {
	return c.startColorFlow(ctx, MethodStartCF, count, action, expressions)
}

// StartBackgroundColorFlow method is used to start a color flow. Color flow is a series of smart LED visible state changing.
// It can be brightness changing, color changing or color temperature changing.
// `count` is the total number of visible state changing before color flow stopped. 0 means infinite loop on the state changing.
func (c Client) StartBackgroundColorFlow(ctx context.Context, count int, action int, expressions []FlowExpression) error {
	return c.startColorFlow(ctx, MethodBgStartCF, count, action, expressions)
}

// StopColorFlow method is used to stop a running color flow.
func (c Client) StopColorFlow(ctx context.Context) error {
	return c.stopColorFlow(ctx, MethodStopCF)
}

// StopBackgroundColorFlow method is used to stop a running color flow.
func (c Client) StopBackgroundColorFlow(ctx context.Context) error {
	return c.stopColorFlow(ctx, MethodBgStopCF)
}

func (c Client) startColorFlow(ctx context.Context, method string, count int, action int, expressions []FlowExpression) error {
	if len(expressions) == 0 {
		return ErrRequiredMinimumOneExpression
	}

	var expressionStr string
	for _, exp := range expressions {
		if len(expressionStr) > 0 {
			expressionStr += ","
		}

		expressionStr += fmt.Sprintf("%d,%d,%d,%d", exp.Duration.Milliseconds(), exp.Mode, exp.Value, exp.Brightness)
	}

	return c.rawWithOk(ctx, method, count, action, expressionStr)
}

func (c Client) stopColorFlow(ctx context.Context, method string) error {
	return c.rawWithOk(ctx, method)
}
