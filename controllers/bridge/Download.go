package bridge

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Minnek-Digital-Studio/cominnek/config"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/folders"
	"github.com/schollz/progressbar/v3"
)

func getRequest(downloadUrl string) *http.Request {
	req, err := http.NewRequest("GET", downloadUrl, nil)

	if err != nil {
		log.Fatalln("ERR1: Someting goes wrong downloading the file")
	}

	return req
}

func download(destinationPath, downloadUrl, fileName string) error {
	tempDestinationPath := destinationPath + ".tmp"

	resp, err := http.DefaultClient.Do(getRequest(downloadUrl))

	if err != nil {
		log.Fatalln("ERR2: Someting goes wrong downloading the file")
	}

	defer resp.Body.Close()

	file, _ := os.OpenFile(tempDestinationPath, os.O_CREATE|os.O_WRONLY, 0644)

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"downloading "+fileName,
	)

	io.Copy(io.MultiWriter(file, bar), resp.Body)
	file.Close()
	os.Rename(tempDestinationPath, destinationPath)

	return nil
}

func DownloadFromURL(url, fileName string) string {
	route := filepath.Join(config.Public.TempPath, fileName)

	if !folders.CheckExists(config.Public.TempPath) {
		folders.Create(config.Public.TempPath)
	}

	download(route, url, fileName)

	println()
	log.Println("🎉 Downloaded")
	return route
}
