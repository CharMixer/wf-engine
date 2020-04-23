package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

/*
backup:
  steps:
    -
      id: fetch_vm_snapshots
      name: "Fetch VM snapshots"
      post_events:
        - fetch_vm_snapshots
      pre_events: []
    -
      id: fetch_data
*/

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
type Step struct {
	Id         string `yaml:"id"`
	Name       string
	PreEvents  []string
	PostEvents []string
}
type Workflow struct {
	Name  string
	Steps []Step `yaml:"steps"`
}

func main() {
	wf := getWorkflow("./wf.yaml")

	fmt.Println(wf)

}

func getWorkflow(path string) Workflow {
	var wf Workflow

	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &wf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	fmt.Printf("%v", wf)

	return wf
}
