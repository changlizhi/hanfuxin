package tests

import (
	"hanfuxin/basemodels"
	"log"
	"testing"
)

func TestMulus(t *testing.T) {
	log.Println(basemodels.Getapppath())
	log.Println(basemodels.Getwenjianmulu("views", "test", "txt"))
	log.Println(basemodels.Getchangliangpath())
	log.Println(basemodels.Getyingyongpath())
}
