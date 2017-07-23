package tests

import (
	"bytes"
	"github.com/astaxie/beego"
	"hanfuxin/basemodels"
	_ "hanfuxin/routers"
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func init() {
	beego.TestBeegoInit(basemodels.Getapppath())
}
func postpatch(bianma string) {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	xx := zfzhi.Xiexianzhi()
	jianhao := zfzhi.Jianhaozhi()
	kzf := zfzhi.Kongzifuzhi()

	reqstr := kzf

	if bianma == zf.POST(false) {
		reqstr = zfzhi.Postjuesezhi()
	}
	if bianma == zf.PATCH(false) {
		reqstr = zfzhi.Patchjuesezhi()
	}
	reqbuf := bytes.NewBuffer([]byte(reqstr))
	req, _ := http.NewRequest(bianma, xx+zf.Juese(true), reqbuf)
	req.Header.Set(zf.Content(false)+jianhao+zf.Type(true), zf.Application(true)+xx+zf.Json(true))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, req)
	log.Println(w.Body.String())

}
func getdelete(bianma string) {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	sz1 := zfzhi.Shuzi1zhi()
	sz1str := strconv.Itoa(sz1)
	xx := zfzhi.Xiexianzhi()

	r, _ := http.NewRequest(bianma, xx+zf.Juese(true)+xx+sz1str, nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	log.Println(w.Body.String())
}
func TestPostjuese(t *testing.T) {
	zf := zf.Zf{}
	postpatch(zf.POST(false))
}
func TestPatchjuese(t *testing.T) {
	zf := zf.Zf{}
	postpatch(zf.PATCH(false))
}
func TestGetjuese(t *testing.T) {
	zf := zf.Zf{}
	getdelete(zf.GET(false))
}
func TestDeletejuese(t *testing.T) {
	zf := zf.Zf{}
	getdelete(zf.DELETE(false))
}
