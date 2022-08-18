package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main()  {

	timeoutContext, cancel := context.WithTimeout(context.Background(), time.Millisecond * 900)
	defer cancel()

	// HTTP 
	//req, err := http.NewRequest(http.MethodGet, "https://via.placeholder.com/350x150", nil)
	req, err := http.NewRequestWithContext(timeoutContext, http.MethodGet, "https://via.placeholder.com/350x150", nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	imageData, err := ioutil.ReadAll(res.Body)

	fmt.Printf("Image is: %d\n", len(imageData))
	
}