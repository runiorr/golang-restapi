package user

type InternalError struct {
	Err error
}

func (i *InternalError) Error() string {
	// return fmt.Sprintf("'Status': '500', 'Error': '%v'", i.Err)
	return "'Status': '500'"
}
