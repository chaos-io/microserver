package handlers

import (
	"io"
	"net/http"
	"path/filepath"
	"strconv"

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

func (f *Files) UploadREST(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fn := vars["filename"]

	f.log.Info("handle POST", "id", id, "filename", fn)

	f.saveFile(id, fn, rw, r.Body)
}

func (f *Files) UploadMultipart(rw http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(128 * 1024); err != nil {
		f.log.Error("Bad request", "error", err)
		http.Error(rw, "Expected multipart form data", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	f.log.Info("Process form for id", "id", id)
	if err != nil {
		f.log.Error("Bad request", "error", err)
		http.Error(rw, "Expected expected integer id", http.StatusBadRequest)
		return
	}

	ff, mh, err := r.FormFile("file")
	if err != nil {
		f.log.Error("Bad request", "error", err)
		http.Error(rw, "Expected file", http.StatusBadRequest)
		return
	}

	f.saveFile(r.FormValue("id"), mh.Filename, rw, ff)
}

func (f *Files) saveFile(id, path string, rw http.ResponseWriter, r io.ReadCloser) {
	f.log.Info("save file for product", "id", id, "path", path)

	fp := filepath.Join(id, path)
	if err := f.store.Save(fp, r); err != nil {
		f.log.Error("unable to save file", "error", err)
		http.Error(rw, "Unable to save file", http.StatusInternalServerError)
	}
}
