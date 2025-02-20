package service

import (
	"io"
	"log"
	"net/http"
)

func GetPublicIP() (string, error) {
	url := "https://api64.ipify.org"
	method := "GET"

	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Println("Erro ao criar a requisicao: ", err)
		return "", err
	}

	res, err := client.Do(req)
	if err != nil {
		log.Println("Erro ao fazer a requisicao: ", err)
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Erro ao ler o corpo da resposta: ", err)
		return "", err
	}

	publicIp := string(body)

	return publicIp, nil

}
