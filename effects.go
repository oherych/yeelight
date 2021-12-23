package yeelight

// Affects
const (
	AffectSudden = "sudden"
	AffectSmooth = "smooth"
)

// IsAffect return true if affect name isRaw correct
func IsAffect(name string) bool {
	return name == AffectSudden || name == AffectSmooth
}

// Affects return list of all supported affects
func Affects() []string {
	return []string{AffectSudden, AffectSmooth}
}
