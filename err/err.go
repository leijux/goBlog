package err

import "github.com/pkg/errors"

var (
	//ErrEnableValue enable err
	ErrEnableValue error = errors.New("Set the correct enable value")
	//ErrOpenFile open file err
	ErrOpenFile error = errors.New("open err")
	//ErrStringIsEmpty  empty
	ErrStringIsEmpty error = errors.New("string is empty")
	
)
