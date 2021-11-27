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
	MethodStartCF               = "start_cf"
	MethodStopCF                = "stop_cf"
	MethodSetScene              = "set_scene"
	MethodCronAdd               = "cron_add"
	MethodCronGet               = "cron_get"
	MethodCronDelete            = "cron_del"
	MethodSetAdjust             = "set_adjust"
	MethodSetMusic              = "set_music"
	MethodSetName               = "set_name"
	MethodSetBgRGB              = "bg_set_rgb"
	MethodSetBgHSV              = "bg_set_hsv"
	MethodSetBgColorTemperature = "bg_set_ct_abx"
	MethodBgStartCF             = "bg_start_cf"
	MethodBgStopCF              = "bg_stop_cf"
	MethodBgSetScene            = "bg_set_scene"
	MethodBgSetDefault          = "bg_set_default"
	MethodBgSetPower            = "bg_set_power"
	MethodSetBgBright           = "bg_set_bright"
	MethodSetBgAdjust           = "bg_set_adjust"
	MethodBgToggle              = "bg_toggle"
	MethodDevToggle             = "dev_toggle"
	MethodAdjustBright          = "adjust_bright"
	MethodAdjustCt              = "adjust_ct"
	MethodAdjustColor           = "adjust_color"
	MethodBgAdjustBright        = "bg_adjust_bright"
	MethodBgAdjustCt            = "bg_adjust_ct"
	MethodBgAdjustColor         = "bg_adjust_color"
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
