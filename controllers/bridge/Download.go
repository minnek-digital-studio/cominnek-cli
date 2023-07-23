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
		log.Fatalln("ERR1: Something goes wrong downloading the file")
	}

	return req
}

func download(destinationPath, downloadUrl, fileName string) error {
	tempDestinationPath := destinationPath

	resp, err := http.DefaultClient.Do(getRequest(downloadUrl))

	if err != nil {
		log.Fatalln("ERR2: Something goes wrong downloading the file")
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

func DownloadTempFromURL(url, fileName string) string {
	route := filepath.Join(config.Public.TempPath, fileName+".tmp")

	if !folders.CheckExists(config.Public.TempPath) {
		folders.Create(config.Public.TempPath)
	}

	download(route, url, fileName)

	println()
	log.Println("🎉 Downloaded")
	return route
}

func DownloadFromURL(url, fileName, path string) string {
	route := filepath.Join(path, fileName)

	if !folders.CheckExists(path) {
		folders.Create(path)
	}

	download(route, url, fileName)

	println()
	log.Println("🎉 Downloaded")
	return route
}

func CheckIfExist(url string) bool {
	resp, err := http.DefaultClient.Do(getRequest(url))

	if err != nil {
		log.Fatalln("ERR: Something goes wrong downloading the file")
	}

	defer resp.Body.Close()

	return resp.StatusCode == 200
}
