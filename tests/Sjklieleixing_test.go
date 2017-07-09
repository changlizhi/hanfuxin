package tests

import (
	"hanfuxin/sjk"
	"log"
	"testing"
)

func TestBianmaLeixing(t *testing.T) {
	sjk := sjk.Sjk{}
	log.Println(sjk.Bianmaleixingzhi())
}
