package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// ! PRUDUCED BY Kamil Umut Araz
// Struct for the API response
type IpstackResponse struct {
	IP        string  `json:"ip"`
	Country   string  `json:"country_name"`
	Region    string  `json:"region_name"`
	City      string  `json:"city"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Error     struct {
		Code int    `json:"code"`
		Info string `json:"info"`
	} `json:"error"`
}

func getLocation(ipAddress string, apiKey string) (IpstackResponse, error) {
	url := fmt.Sprintf("http://api.ipstack.com/%s?access_key=%s", ipAddress, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return IpstackResponse{}, err
	}
	defer resp.Body.Close()

	var location IpstackResponse
	err = json.NewDecoder(resp.Body).Decode(&location)
	if err != nil {
		return IpstackResponse{}, err
	}

	if location.Error.Code != 0 {
		return location, fmt.Errorf("API Error: %s", location.Error.Info)
	}

	return location, nil
}

func main() {
	var ipAddress string
	fmt.Print("Lütfen konumunu öğrenmek istediğiniz IP adresini girin: ")
	fmt.Scan(&ipAddress)

	apiKey := "YOUR_API_KEY" // Buraya kendi API anahtarınızı ekleyin

	location, err := getLocation(ipAddress, apiKey)
	if err != nil {
		fmt.Printf("Bir hata oluştu: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("IP Adresi: %s\n", location.IP)
	fmt.Printf("Ülke: %s\n", location.Country)
	fmt.Printf("Bölge: %s\n", location.Region)
	fmt.Printf("Şehir: %s\n", location.City)
	fmt.Printf("Enlem: %f\n", location.Latitude)
	fmt.Printf("Boylam: %f\n", location.Longitude)
}
