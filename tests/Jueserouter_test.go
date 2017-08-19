package tests

import (
	"bytes"
	"github.com/astaxie/beego"
	"hanfuxin/appinits"
	_ "hanfuxin/routers"
	"changliang/zf"
	"changliang/zfzhi"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func init() {
	beego.TestBeegoInit(appinits.Getapppath())
}
func postpatch(bianma string) {
	reqstr := zfzhi.Zhi.Kzf()

	if bianma == zf.Zfs.POST(false) {
		reqstr = zfzhi.Zhi.Postjuese()
	}
	if bianma == zf.Zfs.PATCH(false) {
		reqstr = zfzhi.Zhi.Patchjuese()
	}
	reqbuf := bytes.NewBuffer([]byte(reqstr))
	req, _ := http.NewRequest(bianma, zfzhi.Zhi.Xx()+zf.Zfs.Juese(true), reqbuf)
	req.Header.Set(zf.Zfs.Content(false)+zfzhi.Zhi.Jia()+zf.Zfs.Type(true), zf.Zfs.Application(true)+zfzhi.Zhi.Xx()+zf.Zfs.Json(true))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, req)
	log.Println(w.Body.String())

}
func getdelete(bianma string) {
	sz1 := zfzhi.Zhi.Shuzi1()
	sz1str := strconv.Itoa(sz1)

	r, _ := http.NewRequest(bianma, zfzhi.Zhi.Xx()+zf.Zfs.Juese(true)+zfzhi.Zhi.Xx()+sz1str, nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	log.Println(w.Body.String())
}
func TestPostjueserouter(t *testing.T) {
	postpatch(zf.Zfs.POST(false))
}
func TestPatchjueserouter(t *testing.T) {
	postpatch(zf.Zfs.PATCH(false))
}
func TestGetjueserouter(t *testing.T) {
	getdelete(zf.Zfs.GET(false))
}
func TestDeletejueserouter(t *testing.T) {
	getdelete(zf.Zfs.DELETE(false))
}
