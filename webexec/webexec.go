package webexec

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mitchellh/mapstructure"
)

func ExecuteCall(path string, kind interface{}) (interface{}, error) {
	log.Printf("Executing call towards: %s", path)
	resp, err := getClient().Get(path)
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
	mapstructure.Decode(temp, &kind)
	return kind, nil
}