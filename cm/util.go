package cm

// SliceContains check that a string slice contains an element or not
func SliceContains(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
