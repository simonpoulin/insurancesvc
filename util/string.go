package util

import (
	"encoding/hex"
	"strings"
)

// ConvertToHex ...
func ConvertToHex(s string) string {

	s = strings.TrimSpace(strings.ToLower(s))
	b := []byte(s)
	result := hex.EncodeToString(b)
	return result
}
