package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func initiateproject() string {

	url := "http://localhost:3001/api/v1/project/Initiate"

	payload := strings.NewReader("{\n\"projectId\": \"projectId-1\"\n, \"name\": \"My Research Project\"\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)

}
