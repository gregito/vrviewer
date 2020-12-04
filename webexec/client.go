package webexec

import (
	"crypto/tls"
	"fmt"
	"github.com/gregito/vrviewer/comp/log"
	"net/http"
)

func getClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		DisableKeepAlives: true,
	}
	c := &http.Client{Transport: tr}
	log.Printf("Creating HTTP client for single use: %s", fmt.Sprintf("%+v\n", c))
	return c
}
