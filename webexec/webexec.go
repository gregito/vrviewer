package webexec

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mitchellh/mapstructure"
)

func MeasureExecuteCall(path string, kind interface{}) (interface{}, error, time.Duration) {
	start := time.Now()
	i, err := ExecuteCall(path, kind)
	return i, err, time.Since(start)
}

func ExecuteCall(path string, kind interface{}) (interface{}, error) {
	log.Println("HTTP request call has been requested")
	client := getClient()
	log.Printf("Executing call towards: %s", path)
	resp, err := client.Get(path)
	if err != nil {
		log.Println(err)
		return kind, err
	}
	log.Printf("Call response: %s\n", fmt.Sprintf("%+v\n", resp))
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
