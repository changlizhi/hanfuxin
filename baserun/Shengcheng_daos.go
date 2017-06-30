package baserun

import (
	"bytes"
	"hanfuxin/baseinits"
	"log"
)

func daosbuffer() {
	log.Println("buffers------------")
}
func Shengchengdaos() {
	buffer := bytes.Buffer{}

	biaos := baseinits.Biaos
	for bk, bv := range biaos {
		log.Println(bv, bk)
	}
	buffer.WriteString("1")
	log.Println("buffrer--------------", buffer.String())
}
