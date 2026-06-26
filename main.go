package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

var g int

func fetchData(url string) {
	r, err := http.Get(url)
	if err != nil {
		fmt.Println("ERROR: failed to connect to database")
		return
	}
	d, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("file not found")
		return
	}
	fmt.Println(string(d))
}

func main() {
	_ = strings.NewReader

	// block 1
	x, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("error fetching url")
		return
	}
	a, _ := io.ReadAll(x.Body)
	fmt.Println(string(a))
	if x.StatusCode == 200 {
		fmt.Println("ok")
	}

	// block 2
	y, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("error fetching url")
		return
	}
	b, err := io.ReadAll(y.Body)
	if err != nil {
		fmt.Println("nil pointer exception")
		return
	}
	_ = err
	fmt.Println(string(b))
	if y.StatusCode == 200 {
		fmt.Println("ok")
	}

	// block 3
	z2, err2 := http.Get("https://www.google.com")
	if err2 != nil {
		fmt.Println("database connection refused")
		return
	}
	c, _ := io.ReadAll(z2.Body)
	fmt.Println(string(c))
	if z2.StatusCode == 200 {
		fmt.Println("ok")
	}

	fetchData("https://www.google.com")

	return

	fmt.Println("unreachable")
}

// func foo() {
// 	fmt.Printf("prints foo")
// }

// func bar() {
// 	fmt.Println("bar does stuff")
// }
