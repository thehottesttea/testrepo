package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UWU", r.URL.Path[1:])
}

/// another empty
//////
func getsomething()string{
	return 	"superkey"
}

func getkey()string{
	password := "superkey"
	return password
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	//dir+"/httpResp/unsubHTML/unsubscribed.png"
	fmt.Println(dir)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		t, err := template.ParseFiles(dir + "/httpResp/unsubHTML/index.html")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			fmt.Println("here1")
			return
		}
		data := map[string]interface{}{"Image": writeImageWithTemplate(dir)}
		err = t.Execute(writer, data)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			fmt.Println("here2")
			return
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func writeImageWithTemplate(dir string) string {
	img, err := ioutil.ReadFile(filepath.Join(dir + "/httpResp/unsubHTML/unsubscribed.png"))
	if err != nil {
		return ""
	}
	originalImage, _, err := image.Decode(bytes.NewReader(img))
	if err != nil {
		return ""
	}
	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, originalImage); err != nil {
		log.Fatalln("unable to encode image.")
	}

	return base64.StdEncoding.EncodeToString(buffer.Bytes())
}
