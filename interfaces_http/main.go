package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type customWriter struct{}

func (customWriter) Write(p []byte) (int, error){
	fmt.Println(string(p))

	return len(p), nil //no errors
}

func main(){
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cw := customWriter{}
	io.Copy(cw, resp.Body)
}
