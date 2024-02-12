package json

import "log"

type CustomError struct {
	message string
	// stack []string
}

func (e CustomError) Error() string {
	return e.message
}

func NewError(message string) error {
	e := CustomError{
		message: message,
	}
	return e
}

func logErrorAndFail(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
