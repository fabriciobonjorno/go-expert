package main

import (
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close() // defer atrasa o processo de execução
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	println(string(res))
}
