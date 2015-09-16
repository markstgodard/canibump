package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const canIBumpUrl = "https://canibump.cfapps.io/"

type status struct {
	CanIBump bool `json:"can_I_bump"`
}

func check(url string) bool {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false
	}

	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	s := &status{}
	err = decoder.Decode(s)
	if err != nil {
		return false
	}

	return s.CanIBump
}

func main() {
	ok := check(canIBumpUrl)
	if ok {
		fmt.Println("Can I bump status is: NO")
		os.Exit(1)
	}
	fmt.Println("Can I bump status is: YES")
}
