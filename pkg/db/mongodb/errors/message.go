package errors

type Err string

func (e Err) Error() string { return string(e) }

const (
	InvalidReturnFromDB = Err("invalid value return from database")
	NoExist             = Err("No Exist")
	Undefined           = Err("Undefined error")
	NilDetect           = Err("Nil Detected")
)
