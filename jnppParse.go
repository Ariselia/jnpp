package jnpp

import (
	"os"
	"path/filepath"

	"github.com/bitly/go-simplejson"
)

/*Parse Parse jspp file*/
func (jn *Jnpp) Parse(script string) error {
	dirn := jn.basedir
	fileloc := dirn + "/" + script
	findex, err := os.Open(fileloc)
	if err != nil {
		return err
	}
	jn.jn, err = simplejson.NewFromReader(findex)
	if err != nil {
		return err
	}

	l, e := filepath.Abs(fileloc)
	if e != nil {
		return e
	}

	jn.basedir = filepath.Dir(l)

	jn.readenv()
	jn.exec()
	jn.merge()
	jn.include()
	jn.render()

	return nil
}
