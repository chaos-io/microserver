package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"

	"chaos-io/microserver/product-images/files"
)

// Files is a handler for reading and writing files
type Files struct {
	log   hclog.Logger
	store files.Storage
}

func NewFiles(s files.Storage, l hclog.Logger) *Files {
	return &Files{log: l, store: s}
}

func (f *Files) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fn := vars["filename"]

	f.log.Info("handle POST", "id", id, "filename", fn)

	f.saveFile(id, fn, rw, r)
}

func (f *Files) saveFile(id, path string, rw http.ResponseWriter, r *http.Request) {
	f.log.Info("save file for product", "id", id, "path", path)

	fp := filepath.Join(id, path)
	if err := f.store.Save(fp, r.Body); err != nil {
		f.log.Error("unable to save file", "error", err)
		http.Error(rw, "Unable to save file", http.StatusInternalServerError)
	}
}
