package kerror

import (
	"fmt"
)

func Wrap(be, err error) error {
	if be == nil {
		return err
	}
	return fmt.Errorf("%s\n%w", be.Error(), err)
}
