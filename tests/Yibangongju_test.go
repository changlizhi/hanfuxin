package tests

import (
	"hanfuxin/apputils"
	"log"
	"testing"
)

func TestPipei(t *testing.T) {
	str := "abc"
	strarr := []string{"abcd", "def", "pfv"}
	log.Println(apputils.Pipei3lei(str, strarr))
	mpall := map[string]string{"abc": "123", "456": "def"}
	log.Println(apputils.Pipei3lei(str, mpall))
}
