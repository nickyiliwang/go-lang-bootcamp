package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Result struct {
	Body string
}

var myClient = &http.Client{Timeout: 10 * time.Second}

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

	// words := strings.Fields(corpus)

	// [1:] get everything after index 0
	query := os.Args[1:]

	fmt.Println(query)
}
