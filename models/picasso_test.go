package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestCategory(t *testing.T) {
	url := "http://service.picasso.adesk.com/v1/vertical/category?adult=false&first=1"
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Host", "service.picasso.adesk.com")
	req.Header.Set("User-Agent", "picasso,268,tencent")
	//req.Header.Set("Accept-Encoding", "gzip")
	req.Header.Set("X-Transmission-Session-Id", " ")

	c := http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%#v\n", resp.Header)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		resp.Body.Close()
		t.Fatalf("%s\n", err)
	}
	resp.Body.Close()
	fmt.Printf("response: %s\n", data)

	message := &Message{}
	if err := json.Unmarshal(data, message); err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("message: %#v\n", message)

	categories := &Categories{}
	if err := json.Unmarshal(message.Res, categories); err != nil {
		t.Fatalf("%s\n", err)
	}
	for _, v := range categories.Data {
		fmt.Printf("%#v\n", *v)
	}
}
