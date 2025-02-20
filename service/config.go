package service

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Configuration struct {
	ZoneID      string
	DNSRecordID string
	Email       string
	APIKEY      string
	Name        string
	URL         string
	SubDomains  string
	Time        string
}

func Config() (c Configuration) {

	rootDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Não foi possível ler o arquivo .env: ", err)
	}

	envFile := filepath.Join(rootDir, ".env")

	err = godotenv.Load(envFile)
	if err != nil {
		log.Fatal("Não foi possível carregar o arquivo .env: ", err)
	}

	c.ZoneID = os.Getenv("ZONE_ID")
	c.DNSRecordID = os.Getenv("DNS_RECORD_ID")
	c.Email = os.Getenv("EMAIL")
	c.APIKEY = os.Getenv("API_KEY")
	c.Name = os.Getenv("NAME")
	c.URL = os.Getenv("URL")
	c.SubDomains = os.Getenv("SUBDOMAINS")
	c.Time = os.Getenv("TIME")

	return c
}
