package jnpp

import (
	"path/filepath"
	"strings"

	"github.com/xiaokangwang/jnpp/jnpputil"
)

func (jn *Jnpp) include() error {
	return jnpputil.FindIf("#!Include", jn.jn, func(fi []string, i interface{}) error {
		return jn.includeat(fi, i)
	}, nil)
}
func (jn *Jnpp) includeat(fi []string, i interface{}) error {
	currentname := fi[len(fi)-1]
	nextname := strings.Split(currentname, "|")[1]
	merging := jn.jn.GetPath(fi...).MustString(nextname + ".json")

	incJn := new(Jnpp)
	incJn.basedir = filepath.Dir(jn.basedir + "/" + merging)
	incJn.environment = jn.environment
	//spew.Dump(merging, incJn)
	err := incJn.Parse(merging)
	if err != nil {
		return err
	}

	cw := jn.jn.GetPath(fi[:len(fi)-1]...)
	cw.Set(nextname, incJn.jn.MustMap())

	cw.Del(currentname)
	return nil
}
