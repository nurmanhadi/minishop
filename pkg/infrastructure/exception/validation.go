package exception

type ErrorValidation struct {
	Message string
}

func (v *ErrorValidation) Error() string {
	return v.Message
}
