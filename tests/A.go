package tests
import(
	"net/http"
)
type Mywriter struct{
}

func (w *Mywriter) Header() http.Header {
	ret := http.Header{}
	return ret
}
func (w *Mywriter) Write([]byte) (int, error) {
	return 1, nil
}
func (w *Mywriter) WriteHeader(par int) {}
