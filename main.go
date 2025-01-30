package main

import (
    "fmt"
    "io"
    "net/http"
)

func main() {

    req, err := http.NewRequest("GET", "https://sandbox-api.dexcom.com/v3/users/self/egvs", nil)
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

    body, _ := io.ReadAll(resp.Body)
    fmt.Println(string(body))

}