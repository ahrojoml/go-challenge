package tickets

import "fmt"

type NullTicketsError struct {
	msg string
}

func (e NullTicketsError) Error() string {
	return fmt.Sprintf("Error: %s", e.msg)
}

func NewNullTicketsError(msg string) NullTicketsError {
	return NullTicketsError{msg: msg}
}

type CountryNotFoundError struct {
	msg string
}

func (e CountryNotFoundError) Error() string {
	return fmt.Sprintf("Error: %s", e.msg)
}

func NewCountryNotFoundError(msg string) CountryNotFoundError {
	return CountryNotFoundError{msg: msg}
}

type InvalidHourError struct{}

func (e InvalidHourError) Error() string {
	return fmt.Sprintf("Error: invalid hour")
}

func NewInvalidHourError() InvalidHourError {
	return InvalidHourError{}
}
