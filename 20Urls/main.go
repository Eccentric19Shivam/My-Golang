package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://offers.hubspot.com/coding-templates?hubs_post-cta=anchor&hubs_post=blog.hubspot.com%2Fwebsite%2Fwebsite-development&hubs_signup-url=blog.hubspot.com%2Fwebsite%2Fwebsite-development&hubs_signup-cta=cta_button"

func main() {
	// fmt.Println("Welcome to handling url in golang")
	// fmt.Println(myUrl)

	//parsing url: getting some specific information from web
	result, err := url.Parse(myUrl)
	checkError(err)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams["hubspot"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	partsOfUrl := &url.URL{ //always pass on the reference-'&'
		Scheme:   "https",
		Host:     "offers.hubspot.com",
		Path:     "/coding-templates",
		RawQuery: "hubs_post-cta=anchor&hubs_post=blog.hubspot.com%2Fwebsite%2Fwebsite-development&hubs_signup-url=blog.hubspot.com%2Fwebsite%2Fwebsite-development&hubs_signup-cta=cta_button",
	}

	anotherURL := partsOfUrl.String() // ye data byte mei tha but .String() krke string mei convert kiya hai
	fmt.Println(anotherURL)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
