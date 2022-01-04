package yeelight

// Effect is enum of Effects
type Effect string

// Effects
const (
	EffectSudden Effect = "sudden"
	EffectSmooth Effect = "smooth"
)

// FlowMode is enum of Flow Mode
type FlowMode int

// Color flow modes
const (
	FlowModeColor        FlowMode = 1
	FlowColorTemperature FlowMode = 2
	FlowSleep            FlowMode = 7
)

// FlowAction is enum of Flow Action
type FlowAction int

// Color flow action
const (
	// FlowActionRecover means smart LED recover to the state before the color flow started.
	FlowActionRecover FlowAction = 0

	// FlowActionStay means smart LED stay at the state when the flow is stopped.
	FlowActionStay FlowAction = 1

	// FlowActionTurnOff means turn off the smart LED after the flow is stopped.
	FlowActionTurnOff FlowAction = 2
)

// AdjustAction is enum of Adjust Actions
type AdjustAction string

// Adjust Actions options
const (
	AdjustActionIncrease AdjustAction = "increase"
	AdjustActionDecrease AdjustAction = "decrease"
	AdjustActionCircle   AdjustAction = "circle"
)

// AdjustProp is enum of Adjust Properties
type AdjustProp string

// Adjust Properties options
const (
	AdjustPropBright           AdjustProp = "bright"
	AdjustPropColorTemperature AdjustProp = "ct"
	AdjustPropColor            AdjustProp = "color"
)
