package main

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
)

func fetchDNSRecords(apiURL, email, apiKey string) ([]DNSRecord, error) {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Auth-Email", email).
		SetHeader("X-Auth-Key", apiKey).
		Get(apiURL)

	if err != nil {
		return nil, err
	}

	var response DNSResponse
	err = json.Unmarshal(resp.Body(), &response)
	if err != nil {
		return nil, err
	}

	return response.Result, nil
}
