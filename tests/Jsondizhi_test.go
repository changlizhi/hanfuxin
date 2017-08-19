package tests

import (
	"hanfuxin/appinits"
	"log"
	"testing"
	"gongju"
)

func TestMulus(t *testing.T) {
	log.Println("apppath-------", appinits.Getapppath())
	log.Println("wenjianmulu------", gongju.Getwenjianmulu("hanfuxin","views", "test", "txt"))
}
