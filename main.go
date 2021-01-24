package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {

	var url string
	fmt.Print("Write: ")
	fmt.Scanf("%s\n", &url)

	del_url := strings.Split(url, "/")

	var file_name string

	if del_url[len(del_url)-1] == ""{
		file_name = del_url[len(del_url)-2]
	} else {
		file_name = del_url[len(del_url)-1]
	}


	err := DownloadFile(file_name, url)
	if err != nil {
		panic(err)
	}
}

func DownloadFile(filepath string, url string) error {

	if !strings.Contains(filepath, ".") {
		filepath += ".html"
	}

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()


	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var buf bytes.Buffer
	endFile := false
	r := io.TeeReader(resp.Body, &buf)
	fmt.Println("Begin download. Already installed:")
	go func() {
		for !endFile {
			fmt.Println(buf.Len()/1024, "Kb")
			time.Sleep(time.Second)
		}
	}()
	_, err = io.Copy(out, r)
	endFile = true

	fmt.Println(buf.Len()/1024, "Kb")
	fmt.Println("Download as " + filepath)
	return err
}