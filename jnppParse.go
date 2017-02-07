package jnpp

import (
	"os"

	"github.com/bitly/go-simplejson"
)

/*Parse Parse jspp file*/
func (jn *Jnpp) Parse(script string) error {
	dirn := jn.basedir
	findex, err := os.Open(dirn + "/" + script)
	if err != nil {
		return err
	}
	jn.jn, err = simplejson.NewFromReader(findex)
	if err != nil {
		return err
	}

	jn.readenv()
	jn.exec()
	jn.merge()
	jn.include()
	jn.render()

	return nil
}
