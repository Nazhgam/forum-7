package cerror

type CError struct {
	Msg string `json:"string"`
}

func (c CError) Error() string {
	return c.Msg
}

func NewError(s string) CError {
	return CError{Msg: s}
}
