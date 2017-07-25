package tests

import (
	"hanfuxin/basemodels"
	"log"
	"testing"
)

func TestMulus(t *testing.T) {
	log.Println("apppath-------", basemodels.Getapppath())
	log.Println("wenjianmulu------", basemodels.Getwenjianmulu("views", "test", "txt"))
}
