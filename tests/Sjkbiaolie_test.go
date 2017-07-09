package tests

import (
	"hanfuxin/sjk"
	"log"
	"testing"
)

func TestYanzhengZhi(t *testing.T) {
	sjk := sjk.Sjk{}
	log.Println(sjk.JuesequanxianJuesebianma(false))
}
