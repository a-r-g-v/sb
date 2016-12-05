package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

var Success int
var Failure int

const (
	maxConnsLimit = 100000
)

var client = http.Client{Transport: &http.Transport{MaxIdleConnsPerHost: maxConnsLimit}}

func TestScenario() {

	resp, _ := client.Get(url)
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	if strings.Contains(string(bytes), "blog") {
		Success++
	} else {
		Failure++
	}
	Success++

}
