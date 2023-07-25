package main

import (
	"fmt"
)

type DNSRecord struct {
	ID       string `json:"id"`
	ZoneName string `json:"zone_name"`
	Type     string `json:"type"`
}

type DNSResponse struct {
	Result []DNSRecord `json:"result"`
}

func main() {
	apiURL := "https://api.cloudflare.com/client/v4/zones/1532b170b4a07c187d0e09c96b8c0567/dns_records"

	email := "nurulyakin79@gmail.com"
	apiKey := "value"

	records, err := fetchDNSRecords(apiURL, email, apiKey) // Menggunakan fungsi fetchDNSRecords dari file dns.go
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, record := range records {
		fmt.Println("id:", record.ID)
		fmt.Println("zone_name:", record.ZoneName)
		fmt.Println("type:", record.Type)
	}
}
