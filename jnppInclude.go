package jnpp

import (
	"strings"

	"github.com/xiaokangwang/jnpp/jnpputil"
)

func (jn *Jnpp) include() {
	jnpputil.FindIf("#!include", jn.jn, func(fi []string, i interface{}) error {
		currentname := fi[len(fi)-1]
		nextname := strings.Split(currentname, "|")[1]
		merging := jn.jn.GetPath(fi...).MustString(nextname + ".json")

		incJn := new(Jnpp)
		incJn.basedir = jn.basedir
		incJn.environment = jn.environment
		err := incJn.Parse(merging)
		if err != nil {
			return err
		}

		jn.jn.GetPath(fi[:len(fi)-2]...).Set(nextname, incJn.jn.MustMap())

		jn.jn.GetPath(fi[:len(fi)-2]...).Del(currentname)
		return nil
	}, nil)
}
