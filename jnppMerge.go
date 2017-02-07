package jnpp

import "github.com/xiaokangwang/jnpp/jnpputil"

func (jn *Jnpp) merge() {
	jnpputil.FindIf("#!merge", jn.jn, func(fi []string, i interface{}) error {
		currentname := fi[len(fi)-1]
		//nextname := strings.Split(currentname, "|")[1]
		merging := jn.jn.GetPath(fi...).MustString()

		incJn := new(Jnpp)
		incJn.basedir = jn.basedir
		incJn.environment = jn.environment
		err := incJn.Parse(merging)
		if err != nil {
			return err
		}
		for k, v := range incJn.jn.MustMap() {
			jn.jn.GetPath(fi[:len(fi)-2]...).Set(k, v)
		}
		jn.jn.GetPath(fi[:len(fi)-2]...).Del(currentname)
		return nil
	}, nil)
}
