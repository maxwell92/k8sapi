package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type paths struct {
	Paths []string `json: "paths"`
}

func Get(url string) (body []byte, err error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func resolveToStruct(response []byte) (s *paths, err error) {
	err = json.Unmarshal(response, &s)
	return s, err
}

func main() {
	url := os.Args[2]
	method := os.Args[1]

	fmt.Println(method)
	fmt.Println(url)

	var response []byte
	response, err := Get(url)
	if err != nil {
		panic(err)
		log.Println(err)
	}

	rs, err := resolveToStruct(response)
	if err != nil {
		panic(err)
		log.Println(err)
	}

	for i := 0; i < len(rs.Paths); i++ {
		fmt.Println(rs.Paths[i])
	}
}
