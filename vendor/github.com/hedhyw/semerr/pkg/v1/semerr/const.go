package semerr

// Error is a constant-like string error.
type Error string

func (err Error) Error() string {
	return string(err)
}
