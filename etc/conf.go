package etc

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type UPList struct {
	AIDs []string `json:"aids"`
	URL  string   `json:"url"`
}

func NewUPList(file string) UPList {
	var uplist UPList
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("ioutil.ReadAll err:", err)
	}
	if err := yaml.Unmarshal(bytes, &uplist); err != nil {
		log.Fatal("yaml unmarshal error :", err)
	}
	fmt.Println("UPList is success")
	return uplist
}
