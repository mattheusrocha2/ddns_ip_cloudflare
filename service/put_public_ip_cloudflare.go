package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type Cloudflare struct {
	Comment  string   `json:"comment"`
	Content  string   `json:"content"`
	Name     string   `json:"name"`
	Proxied  bool     `json:"proxied"`
	Settings Settings `json:"settings"`
	Tags     []Tags   `json:"tags"`
	Ttl      int      `json:"ttl"`
	TypeDNS  string   `json:"type"`
}

type Settings struct {
	Ipv4_only bool `json:"ipv4_only"`
	Ipv6_only bool `json:"ipv6_only"`
}

type Tags struct {
	Owner string `json:"owner"`
}

func PutPublicIPCloudflare(publicIp, dnsRecordID, subdomain string) error {
	config := Config()

	zoneID := config.ZoneID

	uri := fmt.Sprintf("%s/zones/%s/dns_records/%s", config.URL, zoneID, dnsRecordID)
	method := "PUT"

	cloudFlare := Cloudflare{
		Comment: "Domain verification record",
		Content: publicIp, //IP Publico a ser atualizado
		Name:    subdomain,
		Proxied: false,
		Ttl:     3600,
		TypeDNS: "A",
	}

	jsonCloudFlare, err := json.Marshal(cloudFlare)
	if err != nil {
		log.Println("Erro ao criar o arquivo json cloudflare: ", err)
		return err
	}

	payload := strings.NewReader(string(jsonCloudFlare))
	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest(method, uri, payload)
	if err != nil {
		log.Println("Erro ao criar a requisicao em put_public_ip_cloudflare: ", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Email", config.Email)
	req.Header.Set("X-Auth-Key", config.APIKEY)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Erro ao fazer a requisicao em put_public_ip_cloudflare: ", err)
		return err
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Erro ao ler o corpo da resposta em put_public_ip_cloudflare: ", err)
		return err
	}

	return nil
}
