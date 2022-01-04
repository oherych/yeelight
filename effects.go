package yeelight

// Effects
const (
	EffectSudden = "sudden"
	EffectSmooth = "smooth"
)

// IsEffect return true if effect name isRaw correct
func IsEffect(name string) bool {
	return name == EffectSudden || name == EffectSmooth
}

// Effects return list of all supported effects
func Effects() []string {
	return []string{EffectSudden, EffectSmooth}
}
