package error_handler

import (

)

type AlreadyExsistUserError struct {
	Message string
}

func (e AlreadyExsistUserError) Error() string {
    return e.Message
}

type InsertError struct {
	Message string
}

func (e InsertError) Error() string {
	return e.Message
}

type SelectError struct {
	Message string
}

func (e SelectError) Error() string {
	return e.Message
}

type DeleteError struct {
	Message string
}

func (e DeleteError) Error() string {
	return e.Message
}

type UpdateError struct {
	Message string
}

func (e UpdateError) Error() string {
	return e.Message
}

type NoRowsError struct {
	Message string
}

func (e NoRowsError) Error() string {
	return e.Message
}