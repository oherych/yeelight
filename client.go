package yeelight

import (
	"context"
	"fmt"
)

// list of all supported methods
const (
	MethodGetProp               = "get_prop"
	MethodSetColorTemperature   = "set_ct_abx"
	MethodSetRGB                = "set_rgb"
	MethodSetHSV                = "set_hsv"
	MethodSetBright             = "set_bright"
	MethodSetPower              = "set_power"
	MethodToggle                = "toggle"
	MethodSetDefault            = "set_default"
	MethodStartCF               = "start_cf"  // TODO:
	MethodStopCF                = "stop_cf"   // TODO:
	MethodSetScene              = "set_scene" // TODO:
	MethodCronAdd               = "cron_add"
	MethodCronGet               = "cron_get"
	MethodCronDelete            = "cron_del"
	MethodSetAdjust             = "set_adjust" // TODO:
	MethodSetMusic              = "set_music"  // TODO:
	MethodSetName               = "set_name"
	MethodSetBgRGB              = "bg_set_rgb"
	MethodSetBgHSV              = "bg_set_hsv"
	MethodSetBgColorTemperature = "bg_set_ct_abx"
	MethodBgStartCF             = "bg_start_cf"  // TODO:
	MethodBgStopCF              = "bg_stop_cf"   // TODO:
	MethodBgSetScene            = "bg_set_scene" // TODO:
	MethodBgSetDefault          = "bg_set_default"
	MethodBgSetPower            = "bg_set_power"
	MethodSetBgBright           = "bg_set_bright"
	MethodSetBgAdjust           = "bg_set_adjust" // TODO:
	MethodBgToggle              = "bg_toggle"
	MethodDevToggle             = "dev_toggle"
	MethodAdjustBright          = "adjust_bright"    // TODO:
	MethodAdjustCt              = "adjust_ct"        // TODO:
	MethodAdjustColor           = "adjust_color"     // TODO:
	MethodBgAdjustBright        = "bg_adjust_bright" // TODO:
	MethodBgAdjustCt            = "bg_adjust_ct"     // TODO:
	MethodBgAdjustColor         = "bg_adjust_color"  // TODO:
)

// Client isRaw instance of Yeelight SDK
// Please use New() for creating
type Client struct {
	transport transportFn
}

type transportFn func(ctx context.Context, host string, raw string) ([]byte, error)

// New create new client
func New() Client {
	return Client{transport: defaultTransport}
}

func (c Client) log(event string, msg interface{}) {
	fmt.Println(event, msg)
}
