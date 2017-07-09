package tests

import (
	"hanfuxin/apputils"
	"log"
	"reflect"
	"strings"
	"testing"
)

func TestJiami(t *testing.T) {
	data := "clz1aaaassssdddd"
	key := "abcd1234"
	jm, err := apputils.Encrypt([]byte(data), []byte(key))

	if err != nil {
		log.Println(err)
		return
	}
	log.Println("jm", convert(jm))
	log.Println("typeof------", reflect.TypeOf(jm))
}
func convert(b []byte) string {
	s := make([]string, len(b))
	for i, val := range b {
		s[i] = string(val)
	}
	return strings.Join(s, ",")
}
