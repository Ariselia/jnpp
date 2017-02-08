package jnpp

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

/*OpenJnpp jnpp file*/
func OpenJnpp(filename string, environment *map[string]string) (*Jnpp, error) {
	ret := &Jnpp{environment: environment}
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	if !strings.HasPrefix(string(data), "v1.") {
		return nil, errors.New("Version Incorrect")
	}
	dirn := filename + ".d"
	res, err := os.Stat(dirn)
	if err != nil {
		return nil, err
	}
	if !res.IsDir() {
		return nil, errors.New("No .d not dir")
	}
	ret.basedir = dirn

	return ret, nil
}
