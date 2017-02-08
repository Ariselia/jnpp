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

	//var err error
	err = jn.readenv()
	if err != nil {
		return err
	}
	err = jn.exec()
	if err != nil {
		return err
	}
	err = jn.merge()
	if err != nil {
		return err
	}
	err = jn.include()
	if err != nil {
		return err
	}
	err = jn.render()
	if err != nil {
		return err
	}

	return nil
}
