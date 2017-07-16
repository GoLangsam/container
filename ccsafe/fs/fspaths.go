package fs

// Validate returns any errors encountered when validating it's elements
func (f FsPathS) Validate() error {
	er := new(Errors)
	var err error
	for _, fp := range f {
		_, err = fp.Stat()
		er.err(err)
	}
	return er.got()
}

// Accessible returns any errors encountered when accessing it's elements
func (f FsPathS) Accessible() error {
	er := new(Errors)
	for _, fp := range f {
		er.err(fp.Accessible())
	}
	return er.got()
}
