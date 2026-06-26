package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

var g int

func main() {
	_ = strings.NewReader

	x, _ := http.Get("https://www.google.com")
	a, _ := io.ReadAll(x.Body)
	fmt.Println(string(a))

	if x.StatusCode == 200 {
		fmt.Println("ok")
	}

	y, _ := http.Get("https://www.google.com")
	b, _ := io.ReadAll(y.Body)
	fmt.Println(string(b))

	if y.StatusCode == 200 {
		fmt.Println("ok")
	}

	return
}
