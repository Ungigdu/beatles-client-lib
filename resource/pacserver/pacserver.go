package pacserver

import (
	"context"
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/giantliao/beatles-client-lib/config"
	"github.com/giantliao/beatles-client-lib/resource/paclist"
	"github.com/giantliao/beatles-client-lib/resource/webfs"
	"log"
	"strconv"

	"net/http"

	"time"
)

var webserver *http.Server

func StartWebDaemon() {

	mux := http.NewServeMux()

	wfs := assetfs.AssetFS{Asset: webfs.Asset, AssetDir: webfs.AssetDir, AssetInfo: webfs.AssetInfo, Prefix: "resource/localweb"}

	mux.Handle("/web/", http.StripPrefix("/web/", http.FileServer(&wfs)))

	fs := assetfs.AssetFS{Asset: paclist.Asset, AssetDir: paclist.AssetDir, AssetInfo: paclist.AssetInfo, Prefix: "resource/script"}

	mux.Handle("/", http.FileServer(&fs))

	addr := ":" + strconv.Itoa(config.GetCBtlc().StreamServerPacPort)

	log.Println("PAC Server Start at", addr)

	webserver = &http.Server{Addr: addr, Handler: mux}

	log.Fatal(webserver.ListenAndServe())

}

//func StartWebDaemon2() {
//	cfg := config.GetCBtlc()
//	if err := paclist.RestoreAssets(cfg.GetScriptPath(), "resource/script"); err != nil {
//		log.Println("restore script file failed", err)
//	}
//
//	mux := http.NewServeMux()
//
//	mux.HandleFunc("/pac", func(w http.ResponseWriter, r *http.Request) {
//		byt, err := ioutil.ReadFile(path.Join(cfg.GetScriptPath(), "resource/script/gfw.js"))
//		if err != nil {
//			panic(err)
//		}
//		w.Write(byt)
//	})
//
//	addr := ":" + strconv.Itoa(config.GetCBtlc().StreamServerPacPort)
//	log.Println("PAC Server Start at", addr)
//	webserver = &http.Server{Addr: addr, Handler: mux}
//	log.Fatal(webserver.ListenAndServe())
//}

func StopWebDaemon() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	webserver.Shutdown(ctx)
}
