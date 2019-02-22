package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error: os.Getwd: %v", err)
	}
	fileInfoList, err := ioutil.ReadDir(currentDir)
	if err != nil {
		log.Fatalf("Error: ioutil.ReadDir: %v", err)
	}

	var yamlFileNames []string
	for _, fi := range fileInfoList {
		ext := filepath.Ext(fi.Name())
		if ext == ".yml" || ext == ".yaml" {
			yamlFileNames = append(yamlFileNames, fi.Name())
		}
	}
	if len(yamlFileNames) == 0 {
		fmt.Println("Yaml file does not exists.")
		os.Exit(0)
	}

	result := map[string]int{}
	for _, n := range yamlFileNames {
		b, err := ioutil.ReadFile(n)
		if err != nil {
			log.Fatalf("Error: ioutil.ReadFile: %v", err)
		}

		yamlStruct := []struct {
			Roles []string `yaml:"roles"`
		}{}

		if err := yaml.Unmarshal(b, &yamlStruct); err != nil {
			fmt.Printf("Warning: yaml.Unmarshal: %v\n", err)
			continue
		}
		for _, y := range yamlStruct {
			for _, r := range y.Roles {
				result[r]++
			}
		}
	}

	for k, v := range result {
		fmt.Printf("%s: %d\n", k, v)
	}
}
