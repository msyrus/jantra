package jantra

// Error represents a jantra error
type Error struct{ s string }

func (e *Error) Error() string {
	return "jantra: " + e.s
}

// NewError returns a new jantra error
func NewError(s string) *Error {
	return &Error{s}
}
