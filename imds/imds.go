package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func getImds(url string) string {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Metadata", "True")

	client := &http.Client{}

	respString := ""

	for i := 0; i < 60; i++ {
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Attempt %v: Failed with %v\n", i, err)
			time.Sleep(time.Second)
		} else if resp.StatusCode != 200 {
			fmt.Printf("Attempt %v: Failed with %v\n", i, resp.StatusCode)
			time.Sleep(time.Second)
		} else {
			defer resp.Body.Close()
			respBody, _ := ioutil.ReadAll(resp.Body)
			respString = string(respBody)
			break
		}
	}

	return respString
}

func main() {
	url := "http://169.254.169.254/metadata/instance/compute/vmId?api-version=2017-04-02&format=text"

	respBody := getImds(url)
	fmt.Println(string(respBody))
}
