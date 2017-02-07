package main

import (
	"bytes"
	"io"
	"os"

	"github.com/xiaokangwang/jnpp"
)

func main() {
	s, err := jnpp.OpenJnpp(os.Args[1], new(map[string]string))
	if err != nil {
		panic(err)
	}
	err = s.Parse("index.json")
	if err != nil {
		panic(err)
	}
	s2, err := s.Result()
	if err != nil {
		panic(err)
	}
	r := bytes.NewReader(s2)
	_, err = io.Copy(os.Stdout, r)
	if err != nil {
		panic(err)
	}

}
