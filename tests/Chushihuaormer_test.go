package tests
import(
	"hanfuxin/appinits"
	"testing"
	"log"
	"reflect"
)
func TestLogormer(t *testing.T){
	log.Println("Hanfuxinormer-------",appinits.Hanfuxinormer.Driver())

	log.Println("ormertype----------",reflect.TypeOf(appinits.Hanfuxinormer))
}
