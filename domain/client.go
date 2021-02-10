package domain

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//NewClient ...
func NewClient(_host, _fsym, _tsym string) *Client {
	return &Client{
		host: _host,
		fsym: _fsym,
		tsym: _tsym,
	}
}

//Client ...
type Client struct {
	host string
	fsym string
	tsym string
}

//Response ...
func (c *Client) Response() string {
	uri := NewURI(c.host, c.fsym, c.tsym)
	resp, err := http.Get(uri.String())
	if err != nil {
		log.Println(fmt.Sprintf("[Client | Response] Error while get request. Error: %v", err.Error()))
		return ""
	}
	if resp.StatusCode != 200 {
		log.Println(fmt.Sprintf("[Client | Response] Unexpected status code: %v", resp.StatusCode))
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(fmt.Sprintf("[Client | Response] Error while reading response body. Error: %v", err.Error()))
		return ""
	}
	return string(body)
}
