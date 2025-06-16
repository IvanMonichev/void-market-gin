package util

import (
	"os"
	"strings"
)

func SubstitutePlaceholders(s string) string {
	for {
		start := strings.Index(s, "{")
		end := strings.Index(s, "}")
		if start == -1 || end == -1 || end < start {
			break
		}
		key := s[start+1 : end]
		val := os.Getenv(key)
		s = strings.Replace(s, "{"+key+"}", val, 1)
	}
	return s
}
