package jnpp

import "path/filepath"

func (jn *Jnpp) joincalc(merging string) (*Jnpp, error) {
	incJn := new(Jnpp)
	incJn.basedir = filepath.Dir(jn.basedir + "/" + merging)
	incJn.environment = jn.environment
	_, fn := filepath.Split(merging)
	err := incJn.Parse(fn)
	if err != nil {
		return nil, err
	}
	return incJn, nil
}
