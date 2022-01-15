package semerr

import "errors"

// IsTemporaryError checks that error has Temporary method and it
// returns true.
func IsTemporaryError(err error) bool {
	var errTmp interface {
		Temporary() bool
	}

	if errors.As(err, &errTmp) {
		return errTmp.Temporary()
	}

	return false
}
