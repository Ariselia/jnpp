package jnpp

import "github.com/xiaokangwang/jnpp/jnpputil"

func (jn *Jnpp) merge() error {
	return jnpputil.FindIf("#!Merge", jn.jn, func(fi []string, i interface{}) error {
		return jn.mergeat(fi, i)
	}, nil)
}

func (jn *Jnpp) mergeat(fi []string, i interface{}) error {
	currentname := fi[len(fi)-1]
	//nextname := strings.Split(currentname, "|")[1]
	merging := jn.jn.GetPath(fi...).MustString()

	incJn, err := jn.joincalc(merging)
	if err != nil {
		return err
	}

	cw := jn.jn.GetPath(fi[:len(fi)-1]...)

	for k, v := range incJn.jn.MustMap() {
		cw.Set(k, v)
	}
	cw.Del(currentname)
	return nil
}
