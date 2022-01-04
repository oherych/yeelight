package yeelight

// Properties
const (
	PropertyPower        = "power"
	PropertyBright       = "bright"
	PropertyCT           = "ct"
	PropertyRGB          = "rgb"
	PropertyHue          = "hue"
	PropertySat          = "sat"
	PropertyColorMode    = "color_mode"
	PropertyFlowing      = "flowing"
	PropertyDelayOff     = "delayoff"
	PropertyFlowParams   = "flow_params"
	PropertyMusicOn      = "music_on"
	PropertyName         = "name"
	PropertyBgPower      = "bg_power"
	PropertyBgFlowing    = "bg_flowing"
	PropertyBgFlowParams = "bg_flow_params"
	PropertyBgCt         = "bg_ct"
	PropertyBgLMode      = "bg_lmode" // TODO: is this property correct?
	PropertyBgBright     = "bg_bright"
	PropertyBgRgb        = "bg_rgb"
	PropertyBgHue        = "bg_hue"
	PropertyBgSat        = "bg_sat"
	PropertyNlBr         = "nl_br"
	PropertyActiveMode   = "active_mode"
)

// Properties return a list of all susceptible properties
func Properties() []string {
	return []string{
		PropertyPower, PropertyBright, PropertyCT, PropertyRGB, PropertyHue, PropertySat, PropertyColorMode,
		PropertyFlowing, PropertyDelayOff, PropertyFlowParams, PropertyMusicOn, PropertyName, PropertyBgPower,
		PropertyBgFlowing, PropertyBgFlowParams, PropertyBgCt, PropertyBgLMode, PropertyBgBright, PropertyBgRgb,
		PropertyBgHue, PropertyBgSat, PropertyNlBr, PropertyActiveMode,
	}
}
