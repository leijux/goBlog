package err

import "errors"

var (
	ErrEnableValue error = errors.New("Set the correct enable value")
)
