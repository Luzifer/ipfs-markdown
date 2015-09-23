package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Luzifer/rconfig"
	"github.com/flosch/pongo2"
	"github.com/gorilla/mux"
)

//go:generate go-bindata -o bindata.go viewer.html

type config struct {
	IPFSGateway string `flag:"gateway" default:"http://gateway.ipfs.io" description:"Gateway to fetch markdown from"`
	Listen      string `flag:"listen" default:":3000" description:"IP/Port to listen on"`
}

var (
	cfg     = &config{}
	version = "dev"
)

func main() {
	rconfig.Parse(cfg)

	r := mux.NewRouter()
	r.HandleFunc("/document/{ipfs_hash:Qm.+}", handleDocument)
	r.HandleFunc("/{ipfs_hash:Qm.+}", handleViewer)

	http.Handle("/", r)
	http.ListenAndServe(cfg.Listen, r)
}

func handleViewer(res http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	tplData, _ := Asset("viewer.html")
	tpl := pongo2.Must(pongo2.FromString(string(tplData)))
	tpl.ExecuteWriter(pongo2.Context{
		"ipfs_hash": vars["ipfs_hash"],
	}, res)
}

func handleDocument(res http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	documentURL := fmt.Sprintf("%s/ipfs/%s", cfg.IPFSGateway, vars["ipfs_hash"])

	resp, err := http.Get(documentURL)
	if err != nil {
		http.Error(res, "An error ocurred", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	res.Header().Set("Content-Type", "text/markdown")
	io.Copy(res, resp.Body)
}
