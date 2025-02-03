package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	baseURL := "https://sandbox-api.dexcom.com/v3/users/self/egvs"

	params := url.Values{}
	params.Add("startDate", "2022-02-06T09:12:35")
	params.Add("endDate", "2022-02-06T09:12:35")

	fullURL := baseURL + "?" + params.Encode()

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Authorization", "Bearer ZEmdyL0fGCYsC6vC")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Print status code
	fmt.Println("Status Code:", resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: Received non-200 response")
		body, _ := io.ReadAll(resp.Body)
		fmt.Println("Response:", string(body))
		return
	}

	// Read and print response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("Response Body:", string(body))
}