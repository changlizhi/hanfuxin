package tests

import (
	"hanfuxin/appinits"
	"log"
	"reflect"
	"testing"
)

func TestLogormer(t *testing.T) {
	log.Println("Hanfuxinormer-------", appinits.Hanfuxinormer.Driver())

	log.Println("ormertype----------", reflect.TypeOf(appinits.Hanfuxinormer))
}
