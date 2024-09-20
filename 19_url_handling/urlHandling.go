package main

import (
	"fmt"
	"net/url"
)


const MyUrl = "https://go-chi.io/#/README?course=golang"

func main() {
	fmt.Println("Handling URL's!!!")
	result, _ := url.Parse(MyUrl)
	fmt.Println("Scheme is ",result.Scheme)
	fmt.Println("Host is ",result.Host)
	fmt.Println("Path is ",result.Path)
	fmt.Println("Port is ",result.Port())
	fmt.Println("Query Params are ",result.RawQuery)
	qparams := result.Query()
	fmt.Printf("Type of query params using Query() is %T\n", qparams)

	fmt.Println("Course is ", qparams["course"])

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "go-chi.io",
		Path: "/#/README",
		RawPath: "course=golang",
	}

	constructedUrl := partsOfUrl.String()
	fmt.Println(constructedUrl)
}