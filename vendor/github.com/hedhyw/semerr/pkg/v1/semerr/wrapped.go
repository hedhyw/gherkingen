package semerr

// wrappedError wraps error for adding meaning to it.
type wrappedError struct {
	Err error
}

func (err wrappedError) Unwrap() error {
	return err.Err
}

func (err wrappedError) Error() string {
	return err.Err.Error()
}

type temporaryWrappedError struct {
	wrappedError
}

func (err temporaryWrappedError) Temporary() bool {
	return true
}

func newTemporaryWrappedError(err error) temporaryWrappedError {
	return temporaryWrappedError{
		wrappedError: wrappedError{
			Err: err,
		},
	}
}

type permanentWrappedError struct {
	wrappedError
}

func (err permanentWrappedError) Temporary() bool {
	return false
}

func newPermanentWrappedError(err error) permanentWrappedError {
	return permanentWrappedError{
		wrappedError: wrappedError{
			Err: err,
		},
	}
}
