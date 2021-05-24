package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
	//"gopkg.in/yaml.v3"
)

type myData struct {
	Conf struct {
		user  string
		id    string
		empId string `yaml:"empId"`
	}
}

func readConf(filename string) (*myData, error) {
	buf, err := ioutil.ReadFile(config.yaml)
	if err != nil {
		return nil, err
	}

	c := &myData{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}

	return c, nil
}

func main() {
	c, err := readConf("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", c)
}
