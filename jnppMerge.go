package jnpp

import (
	"io/ioutil"
	"path/filepath"

	"github.com/xiaokangwang/jnpp/jnpputil"
)

func (jn *Jnpp) merge() error {
	//Single Merge
	var err error
	err = jnpputil.FindIf("#!MergeAll", jn.jn, func(fi []string, i interface{}) error {
		return jn.mergeat(fi, i, true)
	}, nil)
	if err != nil {
		return err
	}
	err = jnpputil.FindIf("#!Merge", jn.jn, func(fi []string, i interface{}) error {
		return jn.mergeat(fi, i, false)
	}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (jn *Jnpp) mergeat(fi []string, i interface{}, dir bool) error {
	currentname := fi[len(fi)-1]
	cw := jn.jn.GetPath(fi[:len(fi)-1]...)
	//nextname := strings.Split(currentname, "|")[1]
	mergeFile := func(merging string) error {
		incJn, err := jn.joincalc(merging)
		if err != nil {
			return err
		}

		for k, v := range incJn.jn.MustMap() {
			cw.Set(k, v)
		}
		return nil
	}
	merging := jn.jn.GetPath(fi...).MustString()
	if !dir {
		err := mergeFile(merging)
		if err != nil {
			return err
		}
	} else {
		files, err := ioutil.ReadDir(jn.basedir + "/" + merging)
		if err != nil {
			return err
		}
		for _, f := range files {
			fname := f.Name()
			if filepath.Ext(fname) == ".json" {
				err := mergeFile(merging + "/" + fname)
				if err != nil {
					return err
				}
			}
		}

	}

	cw.Del(currentname)
	//spew.Dump(currentname)
	return nil
}
