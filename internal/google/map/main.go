package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const sourcePath = "/Users/bangyuwen/Downloads/dataset_my-task_2020-06-25_06-10-13-203.json"
const sinkPath = "sink.json"

type store struct {
	Title   string
	Reviews []struct {
		Stars int    `json:"stars"`
		Text  string `json:"text"`
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	jsonFile, err := os.Open(sourcePath)
	check(err)
	defer jsonFile.Close()

	os.Remove(sinkPath)
	sinkFile, err := os.OpenFile(sinkPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	defer sinkFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var s store
	json.Unmarshal(byteValue, &s)
	for i, review := range s.Reviews {
		println(i)
		j, _ := json.Marshal(review)
		sinkFile.Write(append(j[:], "\n"...))
	}

}
