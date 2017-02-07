package jnpp

/*Result Parse jspp file*/
func (jn *Jnpp) Result() ([]byte, error) {
	return jn.jn.MarshalJSON()
}
