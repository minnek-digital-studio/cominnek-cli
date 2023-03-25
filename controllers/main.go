package controllers

import (
	"net/http"
	"time"
)

func Connected() (ok bool) {
	timeout := time.Duration(5000 * time.Millisecond)
	client := http.Client{
		Timeout: timeout,
	}

	_, err := client.Get("http://clients3.google.com/generate_204")

	return err == nil
}
