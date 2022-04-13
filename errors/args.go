package errors

// ArgError is struct of argument error
type ArgError struct {
	valueType string
	err       error
}

func NewArgError(err error, vt string) *ArgError {
	return &ArgError{
		valueType: vt,
		err:       err,
	}
}

func (e *ArgError) Error() string {
	return e.err.Error() + " valid value type: " + e.valueType
}
