package models

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type NameIPMatch struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
}

func LoadUsernameSuggestList(path string) (map[string]string, error) {
	jsonData, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	byteData, err := ioutil.ReadAll(jsonData)
	if err != nil {
		return nil, err
	}

	var userIpList []NameIPMatch

	if err := json.Unmarshal(byteData, &userIpList); err != nil {
		return nil, err
	}

	nameIpList := make(map[string]string)
	for _, user := range userIpList {
		nameIpList[user.IP] = user.Name
	}

	return nameIpList, nil
}
