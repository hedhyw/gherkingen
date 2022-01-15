package semerr

import "strings"

// MultiErr is a slice of errors.
type MultiErr []error

func (err MultiErr) Unwrap() error {
	if len(err) == 0 {
		return nil
	}

	return err[0]
}

func (m MultiErr) Error() string {
	strErrs := make([]string, 0, len(m))
	for _, err := range m {
		strErrs = append(strErrs, err.Error())
	}

	return strings.Join(strErrs, "; ")
}

// NewMultiError creates a error that can hold multiple errors.
// It skips or nil values. If count of errors is 1, it returns the
// original value. The main error is the first.
func NewMultiError(errs ...error) error {
	if len(errs) == 0 {
		return nil
	}

	{
		outErrs := make([]error, 0, len(errs))
		for _, err := range errs {
			if err != nil {
				outErrs = append(outErrs, err)
			}
		}
		errs = outErrs
	}

	switch {
	case len(errs) == 0:
		return nil
	case len(errs) == 1:
		return errs[0]
	default:
		return MultiErr(errs)
	}
}
