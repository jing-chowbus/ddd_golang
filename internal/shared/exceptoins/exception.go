package exceptoins

type Exception struct {
	Type    string
	Message string
}

func (exception Exception) Error() string {
	return exception.Type
}

type SystemException struct {
	Exception
}

type NagetiveBalance struct {
	Exception
}
