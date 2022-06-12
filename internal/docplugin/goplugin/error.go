package goplugin

import "fmt"

type outOfRangeError struct {
	Len   int
	Index int
}

func (err outOfRangeError) Error() string {
	return fmt.Sprintf("out of range: %d of %d", err.Index, err.Len)
}
