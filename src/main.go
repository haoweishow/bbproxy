package main

import (
	"fmt"
	"net/http"

	logs "bbproxy.haoweishow.com/log"

	_ "net/http/pprof"

	"github.com/elazarl/goproxy"
)

func main() {
	go func() {
		http.ListenAndServe(":6060", nil)
	}()

	defer logs.Logger.Flush()
	logs.Logger.Info("Hello from Seelog!")
	proxy := goproxy.NewProxyHttpServer()

	proxy.OnResponse().DoFunc(
		func(r *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
			if ctx != nil && ctx.Req != nil && ctx.Resp != nil {
				logs.Logger.Infof("%s, %s, %d", ctx.Req.RemoteAddr, ctx.Req.Host, r.StatusCode)
			}
			return r
		})
	fmt.Println("BBProxy is start....")
	logs.Logger.Info("BBProxy is start....")
	http.ListenAndServe(":8080", proxy)
}
