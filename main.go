package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	apiBaseURL = "https://api.cloudflare.com/client/v4/zones"
)

type DNSRecord struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"` //hanya ip
}

func addDNSRecord(zoneID, apiKey string, record DNSRecord) error {
	url := fmt.Sprintf("%s%s/dns_records", apiBaseURL, zoneID)

	requestData := struct {
		Record DNSRecord `json:"dns_record"`
	}{Record: record}

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to add DNS Record: %s", resp.Status)
	}
	fmt.Println("DNS Berhasil ditambahkan untuk zone dengan ID:", zoneID)
	return nil
}

func main() {
	apiKey := "API_KEY"
	zoneID := "ZONE_ID"

	recordToAdd := DNSRecord{
		Type:    "A",                  //type A record
		Name:    "halo.bukaevent.com", //subdomain
		Content: "0.0.0.0",            //IP origin
	}
	err := addDNSRecord(zoneID, apiKey, recordToAdd)
	if err != nil {
		fmt.Println("Error", err)
	}
}
