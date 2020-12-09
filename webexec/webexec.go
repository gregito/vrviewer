package webexec

import (
	"encoding/json"
	"fmt"
	"github.com/gregito/vrviewer/comp/log"
	"net/http"
	"time"

	"github.com/mitchellh/mapstructure"
)

func MeasuredExecuteCallWithClient(client *http.Client, path string, kind interface{}) (interface{}, error, time.Duration) {
	start := time.Now()
	i, err := executeCallWithClient(client, path, kind)
	return i, err, time.Since(start)
}

func executeCallWithClient(client *http.Client, path string, kind interface{}) (interface{}, error) {
	log.Println("HTTP request call has been requested")
	log.Printf("Executing call towards: %s", path)
	resp, err := client.Get(path)
	if err != nil {
		log.Println(err)
		return kind, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Printf("Error! Status code was:%d\n", resp.StatusCode)
		return kind, nil
	}
	var temp interface{}
	_ = json.NewDecoder(resp.Body).Decode(&temp)
	if err = mapstructure.Decode(temp, &kind); err != nil {
		log.Println(fmt.Sprintf("Unable to decode content from map to: %T\n", kind))
	}
	return kind, nil
}
