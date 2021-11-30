package yeelight

import (
	"context"
	"fmt"
)

// list of all supported methods
const (
	MethodGetProp               = "get_prop"
	MethodSetColorTemperature   = "set_ct_abx"       // TODO:
	MethodSetRGB                = "set_rgb"          // TODO:
	MethodSetHSV                = "set_hsv"          // TODO:
	MethodSetBright             = "set_bright"       // TODO:
	MethodSetPower              = "set_power"        // TODO:
	MethodToggle                = "toggle"           // TODO:
	MethodSetDefault            = "set_default"      // TODO:
	MethodStartCF               = "start_cf"         // TODO:
	MethodStopCF                = "stop_cf"          // TODO:
	MethodSetScene              = "set_scene"        // TODO:
	MethodCronAdd               = "cron_add"         // TODO:
	MethodCronGet               = "cron_get"         // TODO:
	MethodCronDelete            = "cron_del"         // TODO:
	MethodSetAdjust             = "set_adjust"       // TODO:
	MethodSetMusic              = "set_music"        // TODO:
	MethodSetName               = "set_name"         // TODO:
	MethodSetBgRGB              = "bg_set_rgb"       // TODO:
	MethodSetBgHSV              = "bg_set_hsv"       // TODO:
	MethodSetBgColorTemperature = "bg_set_ct_abx"    // TODO:
	MethodBgStartCF             = "bg_start_cf"      // TODO:
	MethodBgStopCF              = "bg_stop_cf"       // TODO:
	MethodBgSetScene            = "bg_set_scene"     // TODO:
	MethodBgSetDefault          = "bg_set_default"   // TODO:
	MethodBgSetPower            = "bg_set_power"     // TODO:
	MethodSetBgBright           = "bg_set_bright"    // TODO:
	MethodSetBgAdjust           = "bg_set_adjust"    // TODO:
	MethodBgToggle              = "bg_toggle"        // TODO:
	MethodDevToggle             = "dev_toggle"       // TODO:
	MethodAdjustBright          = "adjust_bright"    // TODO:
	MethodAdjustCt              = "adjust_ct"        // TODO:
	MethodAdjustColor           = "adjust_color"     // TODO:
	MethodBgAdjustBright        = "bg_adjust_bright" // TODO:
	MethodBgAdjustCt            = "bg_adjust_ct"     // TODO:
	MethodBgAdjustColor         = "bg_adjust_color"  // TODO:
)

type Client struct {
	transport transportFn
}

type transportFn func(ctx context.Context, host string, raw string) ([]byte, error)

func New() Client {
	return Client{transport: defaultTransport}
}

func (c Client) log(event string, msg interface{}) {
	fmt.Println(event, msg)
}
