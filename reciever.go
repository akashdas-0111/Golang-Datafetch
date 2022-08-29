package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://go.dev/"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	// defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}
