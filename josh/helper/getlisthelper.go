package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type input struct {
	Website []string `json:"website"`
}

func GetList(r *http.Request) (map[string]bool, error) {
	var userInput = input{}

	var websiteList = make(map[string]bool)

	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for _, val := range userInput.Website {
		websiteList[val] = false
	}
	return websiteList, nil
}
