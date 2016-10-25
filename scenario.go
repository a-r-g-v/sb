package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

var Success int
var Failure int

func TestScenario() {

	url := "https://api.push7.jp/api/v1/"
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	if strings.Contains(string(bytes), "9aea470") {
		Success++
	} else {
		Failure++
	}

}
