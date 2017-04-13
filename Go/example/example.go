package main

import (
	"fmt"

	".."
)

func main() {
	request := new(unu.Request)
	request.URL = "https://xtom.com/" // URL to shrink
	request.Title = "xTom"            // optional, if omitted your url will lookup title
	request.Keyword = "xTom"          // optional keyword
	response, err := unu.Submit(request)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
