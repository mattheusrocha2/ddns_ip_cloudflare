package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type DNSKey struct {
	Result     []Result `json:"result"`
	Sucess     bool     `json:"success"`
	Errors     []error  `json:"errors"`
	Messages   []string `json:"messages"`
	ResultInfo struct {
		Page       int `json:"page"`
		PerPage    int `json:"per_page"`
		TotalCount int `json:"total_count"`
		TotalPages int `json:"total_pages"`
	} `json:"result_info"`
}

type Result struct {
	ID         string   `json:"id"`
	ZoneID     string   `json:"zone_id"`
	ZoneName   string   `json:"zone_name"`
	Name       string   `json:"name"`
	Type       string   `json:"type"`
	Content    string   `json:"content"`
	Proxiable  bool     `json:"proxiable"`
	Proxied    bool     `json:"proxied"`
	TTL        int      `json:"ttl"`
	Settings   struct{} `json:"settings"`
	Meta       struct{} `json:"meta"`
	Comment    string   `json:"comment"` // Usa *string para permitir valores nulos
	Tags       []string `json:"tags"`
	CreatedOn  string   `json:"created_on"`
	ModifiedOn string   `json:"modified_on"`
}

func GetDSNKey() (DNSKey, error) {
	config := Config()

	uri := fmt.Sprintf("%s/zones/%s/dns_records", config.URL, config.ZoneID)
	method := "GET"

	client := &http.Client{}

	req, err := http.NewRequest(method, uri, nil)
	if err != nil {
		log.Println("Erro ao criar a requisicao GetDNSKey: ", err)
		return DNSKey{}, err
	}

	req.Header.Set("X-Auth-Email", Config().Email)
	req.Header.Set("X-Auth-Key", Config().APIKEY)

	res, err := client.Do(req)
	if err != nil {
		log.Println("Erro ao fazer a requisicao GetDNSKey: ", err)
		return DNSKey{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Erro ao ler o corpo da resposta: ", err)
		return DNSKey{}, err
	}

	var bodyDNS DNSKey

	err = json.Unmarshal(body, &bodyDNS)
	if err != nil {
		log.Println("Erro ao deserializar o JSON em GetDNSKey: ", err)
		return DNSKey{}, err
	}

	return bodyDNS, nil

}
