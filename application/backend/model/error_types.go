package model

type Error string
func (e Error) Error() string {
	return string(e)
}

const (
	ErrDataIsNil = Error("Unexpected nil model data.")
)
