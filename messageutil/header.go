// Package messageutil provides MIME utility functions.
//
// It should only be used by libraries working with low-level MIME.
package messageutil

import (
	"strings"
)

const maxHeaderLen = 76

// FormatHeaderField formats a header field, ensuring each line is no longer
// than 76 characters. It tries to fold lines at whitespace characters if
// possible. If the header contains a word longer than this limit, it will be
// split.
func FormatHeaderField(k, v string) string {
	s := k + ": "

	first := true
	for len(v) > 0 {
		maxlen := maxHeaderLen
		if first {
			maxlen -= len(s)
		}

		// We'll need to fold before i
		foldBefore := maxlen + 1
		foldAt := len(v)

		var folding string
		if foldBefore > len(v) {
			// We reached the end of the string
			if v[len(v)-1] != '\n' {
				// If there isn't already a trailing CRLF, insert one
				folding = "\r\n"
			}
		} else {
			// Find the closest whitespace before i
			foldAt = strings.LastIndexAny(v[:foldBefore], " \t\n")
			if foldAt == 0 {
				// The whitespace we found was the previous folding WSP
				foldAt = foldBefore - 1
			} else if foldAt < 0 {
				// We didn't find any whitespace, we have to insert one
				foldAt = foldBefore - 2
			}

			switch v[foldAt] {
			case ' ', '\t':
				if v[foldAt-1] != '\n' {
					folding = "\r\n" // The next char will be a WSP, don't need to insert one
				}
			case '\n':
				folding = "" // There is already a CRLF, nothing to do
			default:
				folding = "\r\n " // Another char, we need to insert CRLF + WSP
			}
		}

		s += v[:foldAt] + folding
		v = v[foldAt:]
		first = false
	}

	return s
}
