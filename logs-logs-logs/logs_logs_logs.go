package logs

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, c := range log {
		switch c {
		case '\u2757':
			return "recommendation"
		case '\U0001f50d':
			return "search"
		case '\u2600':
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	res := ""
	for _, c := range log {
		if c == oldRune {
			res = res + string(newRune)
		} else {
			res = res + string(c)
		}
	}
	return string(res)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return len([]rune(log)) <= limit
}
