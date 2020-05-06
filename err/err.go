package err

import "errors"

var (
	//ErrEnableValue enable err
	ErrEnableValue error = errors.New("Set the correct enable value")
	//ErrOpenFile open file err
	ErrOpenFile error=errors.New("open err:")
)
