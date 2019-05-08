package cm

import (
	"fmt"
	"os"
)

// SliceContains check that a string slice contains an element or not
func SliceContains(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// ByteCountDecimal create human readable format for bytes
func ByteCountDecimal(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
}

// Exists reports whether the named file or directory exists.
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func AddQutes(slice []string) []string {
	newSlice := make([]string, 0)
	for _, str := range slice {
		newSlice = append(newSlice, "\""+str+"\"")
	}
	return newSlice
}
