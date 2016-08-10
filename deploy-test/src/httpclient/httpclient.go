package httpclient

import (
	"crypto/tls"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type HttpClient struct {
	client *http.Client
	host   string
	port   string
}

var instance *HttpClient
var once sync.Once

func HttpInstance() *HttpClient {
	return instance
}

func NewHttpClient(host, port string) *HttpClient {
	once.Do(func() {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // TODO:  tls not used or tls undefined
		}
		client := &http.Client{Transport: tr}
		instance = &HttpClient{client: client, host: host, port: port}
	})
	return instance
}

func (c *HttpClient) Get(url string) (response []byte, err error) {
	resp, err := c.client.Get(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer resp.Body.Close()

	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return response, nil
}
func (c *HttpClient) Post(url string, body io.Reader) (response []byte, err error) {
	resp, err := c.client.Post(url, "application/json;charset=utf-8", body)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer resp.Body.Close()

	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return response, nil
}

//TODO: add Http Put Method
func (c *HttpClient) Put() {

}

func (c *HttpClient) Delete(url string) (response []byte, err error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer resp.Body.Close()

	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return response, nil
}
