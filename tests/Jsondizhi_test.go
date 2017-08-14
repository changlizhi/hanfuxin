package tests

import (
	"hanfuxin/appinits"
	"log"
	"testing"
)

func TestMulus(t *testing.T) {
	log.Println("apppath-------", appinits.Getapppath())
	log.Println("wenjianmulu------", appinits.Getwenjianmulu("views", "test", "txt"))
}
