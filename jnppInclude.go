package jnpp

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/xiaokangwang/jnpp/jnpputil"
)

func (jn *Jnpp) include() error {
	var err error
	err = jnpputil.FindIf("#!IncludeAll", jn.jn, func(fi []string, i interface{}) error {
		return jn.includeat(fi, i, true)
	}, nil)
	if err != nil {
		return err
	}
	err = jnpputil.FindIf("#!Include", jn.jn, func(fi []string, i interface{}) error {
		return jn.includeat(fi, i, false)
	}, nil)
	if err != nil {
		return err
	}
	return nil
}
func (jn *Jnpp) includeat(fi []string, i interface{}, dir bool) error {
	currentname := fi[len(fi)-1]
	nextname := strings.Split(currentname, "|")[1]
	cw := jn.jn.GetPath(fi[:len(fi)-1]...)
	var merging string
	if !dir {
		merging = jn.jn.GetPath(fi...).MustString(nextname + ".json")
		incJn, err := jn.joincalc(merging)
		if err != nil {
			return err
		}

		cw.Set(nextname, incJn.jn.MustMap())
	} else {
		merging = jn.jn.GetPath(fi...).MustString(nextname + ".d")
		files, err := ioutil.ReadDir(jn.basedir + "/" + merging)
		if err != nil {
			return err
		}
		inserting := make([](interface{}), 0, len(files))
		for _, f := range files {
			fname := f.Name()
			if filepath.Ext(fname) == ".json" {
				incJn, err := jn.joincalc(merging + "/" + fname)
				if err != nil {
					return err
				}
				inserting = append(inserting, incJn.jn)
				if err != nil {
					return err
				}
			}
		}
		cw.Set(nextname, inserting)
	}

	cw.Del(currentname)
	return nil
}
