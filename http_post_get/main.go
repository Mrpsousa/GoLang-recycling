package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func post() {
	c := http.Client{}
	jsonVar := bytes.NewBuffer([]byte(`{"name":"Roger"}`))
	req, err := http.NewRequest("POST", "http://google.com", jsonVar)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/json")
	res, err := c.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	io.CopyBuffer(os.Stdout, res.Body, nil)
}

func get() {
	c := http.Client{}
	req, err := http.NewRequest("GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/json")
	res, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func main() {
	post()
	get()
}
