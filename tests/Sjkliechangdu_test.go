package tests

import (
	"hanfuxin/sjk"
	"log"
	"testing"
)

func TestMingchengchangduzhi(t *testing.T) {
	sjk := sjk.Sjk{}
	log.Println(sjk.Mingchengchangdu())
}
