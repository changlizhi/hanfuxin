package tests

import (
	"bytes"
	"github.com/astaxie/beego"
	"hanfuxin/basemodels"
	_ "hanfuxin/routers"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	beego.TestBeegoInit(basemodels.Getapppath())
}
func TestZengjiaget(t *testing.T) {
	r, _ := http.NewRequest("GET", "/juese", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	log.Println(w.Body.String())
}
func TestZengjiapost(t *testing.T) {
	reqstr := `{"Id":4,"Bianma":"ROLE","Mingcheng":"角色1","Biaoji":"Youxiao"}`
	reqbuf := bytes.NewBuffer([]byte(reqstr))
	req, _ := http.NewRequest("POST", "/juese", reqbuf)
	req.Header.Set("Content-type", "application/json")
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, req)
	log.Println(w.Body.String())
}
func TestZengjiahttppost(t *testing.T) {
	reqstr := `{"Id":4,"Bianma":"ROLE","Mingcheng":"角色1","Biaoji":"Youxiao"}`
	reqbuf := bytes.NewBuffer([]byte(reqstr))
	req, _ := http.NewRequest("POST", "http://localhost:8080/juese", reqbuf)
	req.Header.Set("Content-type", "application/json")
	c := &http.Client{}
	res, _ := c.Do(req)
	body, _ := ioutil.ReadAll(res.Body)
	log.Println(string(body))
}
