package main

import (
	"fmt"
	"io"
	_ "log"
	"net/http"
	"os"
)

func HelloServer(w http.ResponseWriter, r *http.Request) {
	if "POST" == r.Method {
		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		fmt.Println(handler.Header)
		defer file.Close()
		f, err := os.OpenFile("./"+handler.Filename,
			os.O_WRONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		size, err := io.Copy(f, file)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Fprintf(w, "上传文件的大小为: %d", size)
		return
	} // 上传页面
	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(200)
	html := `<form enctype="multipart/form-data" action="/" method="POST"> The file you want to upload: <input name="file" type="file" /><br/> <input type="submit" value="Upload File" />
			</form> `
	io.WriteString(w, html)
}
func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		fmt.Println(err)
	}
}
