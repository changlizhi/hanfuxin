package tests

import (
	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
	"hanfuxin/appinits"
	_ "hanfuxin/routers"
	"net/http"
	"net/http/httptest"
	"testing"
	"log"
	"bytes"
	"io/ioutil"
)

func init() {
	beego.TestBeegoInit(appinits.Getapppath())
}

// TestBeego is a sample to run an endpoint test
func TestBeego(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestBeego", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})

		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
func TestZengjiahttppost(t *testing.T) {
     reqstr := `{"Id":9,"Bianma":"ROLE","Mingcheng":"我的角色","Biaoji":"Youxiao"}`
     reqbuf := bytes.NewBuffer([]byte(reqstr))
     req, _ := http.NewRequest("POST", "http://localhost:8080/juese", reqbuf)
     req.Header.Set("Content-type", "application/json")
     c := &http.Client{}
     res, _ := c.Do(req)
     body, _ := ioutil.ReadAll(res.Body)
     log.Println(string(body))
}

