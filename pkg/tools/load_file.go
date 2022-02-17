package tools

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

func LoadFile(w http.ResponseWriter, r *http.Request, id int, dir string) {
	Logger.Error("start function")
	uploadfile, handle, err := r.FormFile("image")
	if ErrorManager(err, w) {
		return
	}
	if ext := strings.ToLower(path.Ext(handle.Filename)); ext != ".jpg" && ext != ".png" {
		http.Error(w, "Only png/jpg file", http.StatusInternalServerError)
		return
	}
	os.Mkdir(fmt.Sprint("./assets/images/", dir), 0777)
	saveFile, err := os.OpenFile(fmt.Sprint("./assets/images/", dir, "/", id, ".png"), os.O_WRONLY|os.O_CREATE, 0666)
	if ErrorManager(err, w) {
		return
	}
	io.Copy(saveFile, uploadfile)
	defer uploadfile.Close()
	defer saveFile.Close()
	return
}
