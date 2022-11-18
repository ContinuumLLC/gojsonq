package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	gojsonq "github.com/thedevsaddam/gojsonq/v2"
)

func main() {
	content, err := ioutil.ReadFile("/home/fusion/workspace/src/github.com/lkumarjain/playground/event_log_array.json")
	//content, err := ioutil.ReadFile("/home/fusion/workspace/src/github.com/lkumarjain/playground/event_log_object.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	result, err := query(content, "RegID AS id,PartnerID,EndpointID,Message.hostname as HostName,test")

	log.Printf("Content %s\n", content)
	log.Printf("Result: %s\n", result)
	log.Printf("Error: %v", err)
}

func query(content []byte, selection string) (string, error) {
	jq := gojsonq.New().FromByteArray(content)

	jq.Where("Action", "=", "eventLog1")
	jq.OrWhere("AgentID", "=", "9ffc24a3-ec95-49d5-8e4e-5ac6a6a9762c")
	jq.Select(strings.Split(selection, ",")...)
	jq.DefaultValues("test", "test")
	var b bytes.Buffer
	jq.Writer(&b)

	fmt.Println("Errors ::: ", jq.Errors())

	return b.String(), jq.Error()
}
