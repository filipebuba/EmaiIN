package internalError

import "errors"

var ErrInternal error = errors.New("internal Server error")
