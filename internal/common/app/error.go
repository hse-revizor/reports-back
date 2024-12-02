package app

import (
	"errors"
)

var InternalErr = errors.New("Internal error")
var InvalidDataErr = errors.New("Bad request")
var NotFoundErr = errors.New("Not found")