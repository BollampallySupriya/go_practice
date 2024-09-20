package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to API request handling tutorial!!")
	// PerformGetRequest()
	// PerformPostJSONRequest()
	PerformPostFormRequest()
}


func PerformGetRequest() {
	const url = "http://localhost:8000/get"
	response, err := http.Get(url)
	HandleNilError(err)
	defer response.Body.Close()

	// fmt.Println("Status Code is ",response.StatusCode)
	// fmt.Println("Content Length is", response.ContentLength)

	var responseString strings.Builder
	byte_data, err := io.ReadAll(response.Body)
	HandleNilError(err)
	fmt.Println(string(byte_data))

	byte_count, err := responseString.Write(byte_data)
	HandleNilError(err)

	
	fmt.Println("byte count is ", byte_count)
	fmt.Println("response string is ", responseString.String())
}

func PerformPostJSONRequest() {
	const myurl = "http://localhost:8000/post"

	requestBody := strings.NewReader(`{
		"coursename": "let's go with golang",
		"price": 4657
	}`)
	response, err := http.Post(myurl, "application/json", requestBody)
	HandleNilError(err)
	defer response.Body.Close()
	byte_data, err := io.ReadAll(response.Body)
	HandleNilError(err)
	fmt.Println(string(byte_data))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("firstname", "Supriya")
	data.Add("lastname", "Bollampally")

	response, err := http.PostForm(myurl, data)
	HandleNilError(err)
	defer response.Body.Close()
	byte_data, err := io.ReadAll(response.Body)
	HandleNilError(err)
	fmt.Println(string(byte_data))

}

func HandleNilError(err error) {
	if err != nil {
		panic(err)
	}
}