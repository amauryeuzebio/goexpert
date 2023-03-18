package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func main() {
	c := http.Client{}

	jsonVar := bytes.NewBuffer([]byte(`{"name": "Amaury"}`))

	resp, err := c.Post("http://google.com", "application/json", jsonVar)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	println(string(body))
}
