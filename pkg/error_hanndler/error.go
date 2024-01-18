package error_hanndler

import (

)

type AlreadyExsistUserError struct {
	Message string
}

func (e AlreadyExsistUserError) Error() string {
    return e.Message
}