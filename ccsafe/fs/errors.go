package fs

const (
	Tab = "\t"
)

type Errors struct {
	errs []error
}

// Errors returns a Tab-terminated string listing the accumulated errors
// (the string is suitable for text\tabwriter;
// if there are no errors, a single Tab is returned - an empty Tab-terminated string)
func (er Errors) Error() string {
	s := Tab
	for _, err := range er.errs {
		s = s + err.Error() + Tab
	}
	return s
}

// err adds any non-nil error
func (er Errors) err(err error) {
	if err != nil {
		er.errs = append(er.errs, err)
	}
}

// got returns the error(s) got, or nil, if there are none
func (er Errors) got() error {
	if len(er.errs) > 0 {
		return er
	} else {
		return nil
	}
}