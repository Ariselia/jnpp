package jnpp

/*
Jnpp is a general purposed json preprocessor.
*/

import "github.com/bitly/go-simplejson"

/*Jnpp Representing a working Jnpp file*/
type Jnpp struct {
	jn          *simplejson.Json
	basedir     string
	environment *map[string]string
}
