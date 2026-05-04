package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ContainerData struct {
	ContainerID          string `json:"container_id"`
	ContainerStatus      string `json:"container_status"`
	ShippedFrom          string `json:"shipped_from"`
	ShippedTo            string `json:"shipped_to"`
	EtaFinalDestination  string `json:"eta_final_destination"`
}

type ContainerResponse struct {
	Data ContainerData `json:"data"`
}

func main() {
	apiKey := os.Getenv("JSONCARGO_API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "JSONCARGO_API_KEY environment variable is not set")
		os.Exit(1)
	}

	trackingNumber := "MSCU1234567" // replace with a real container number

	// If the container prefix is shared across carriers, append ?shipping_line=MSC
	url := fmt.Sprintf("http://api.jsoncargo.com/api/v1/containers/%s", trackingNumber)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to create request:", err)
		os.Exit(1)
	}
	req.Header.Set("x-api-key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Request failed:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to read response:", err)
		os.Exit(1)
	}

	if resp.StatusCode != 200 {
		fmt.Fprintf(os.Stderr, "Error %d: %s\n", resp.StatusCode, string(body))
		os.Exit(1)
	}

	var result ContainerResponse
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to parse response:", err)
		os.Exit(1)
	}

	d := result.Data
	fmt.Println("Container: ", d.ContainerID)
	fmt.Println("Status:    ", d.ContainerStatus)
	fmt.Println("From:      ", d.ShippedFrom)
	fmt.Println("To:        ", d.ShippedTo)
	fmt.Println("ETA:       ", d.EtaFinalDestination)
}
