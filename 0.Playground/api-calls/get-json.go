// https://stackoverflow.com/questions/17156371/how-to-get-json-response-from-http-get
package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type Result struct {
	Body string
}

var myClient = &http.Client{Timeout: 10 * time.Second}

// decoder on the reader directly. Here's a nice function that gets a url and decodes its response onto a target structure.

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func main() {
	res := Result{}
	getJson("https://jsonplaceholder.typicode.com/posts/1", &res)
	println(res.Body)
}
