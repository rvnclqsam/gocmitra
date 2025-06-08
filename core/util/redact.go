package util

// Redact returns a masked version of a sensitive string,
// keeping the first 3 and last 3 characters visible.
func Redact(s string) string {
	if len(s) <= 6 {
		return "***"
	}
	return s[:3] + "..." + s[len(s)-3:]
}
