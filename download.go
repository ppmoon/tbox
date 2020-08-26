package tbox

import (
	"net/http"
	"os"
)

// download file by url
func DownloadFile(url string) (fileByte []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	fileByte, err = IoReaderToByte(resp.Body)
	return
}

// download and store the file path look like /mnt/logo.png
func DownloadFileAndStore(filePath, url string) (err error) {
	fileByte, err := DownloadFile(url)
	if err != nil {
		return
	}
	out, err := os.Create(filePath)
	if err != nil {
		return
	}
	defer out.Close()
	_, err = out.Write(fileByte)
	return
}
