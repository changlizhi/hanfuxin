package tests

import (
	"log"
	"testing"
	"hanfuxin/sjk"
)

func TestYanzhengZhi(t *testing.T) {
	sjk := sjk.Sjk{}
	log.Println(sjk.JuesequanxianJuesebianma(false))
}
