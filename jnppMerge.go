package jnpp

import (
	"path/filepath"

	"github.com/xiaokangwang/jnpp/jnpputil"
)

func (jn *Jnpp) merge() {
	jnpputil.FindIf("#!merge", jn.jn, func(fi []string, i interface{}) error {
		currentname := fi[len(fi)-1]
		//nextname := strings.Split(currentname, "|")[1]
		merging := jn.jn.GetPath(fi...).MustString()

		incJn := new(Jnpp)
		incJn.basedir = filepath.Dir(jn.basedir + "/" + merging)
		incJn.environment = jn.environment
		err := incJn.Parse(merging)
		if err != nil {
			return err
		}

		cw := jn.jn.GetPath(fi[:len(fi)-1]...)

		for k, v := range incJn.jn.MustMap() {
			cw.Set(k, v)
		}
		cw.Del(currentname)
		return nil
	}, nil)
}
