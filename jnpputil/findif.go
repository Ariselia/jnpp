package jnpputil

import (
	"strings"

	"github.com/bitly/go-simplejson"
)

/*FindIf ...*/
func FindIf(target string, within *simplejson.Json, cb func([]string, interface{}) error, cba interface{}) error {
	return findIf(target, within, cb, cba, make([]string, 0, 15))
}

func findIf(target string, within *simplejson.Json, cb func([]string, interface{}) error, cba interface{}, track []string) error {
	for s := range within.MustMap() {
		if strings.HasPrefix(s, target) {
			err := cb(append(track, s), cba)
			if err != nil {
				return err
			}
		} else {
			_, err := within.Get(s).Map()
			if err != nil {
				continue
			}
			err = findIf(target, within.Get(s), cb, cba, append(track, s))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
