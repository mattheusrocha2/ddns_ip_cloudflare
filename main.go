package main

import (
	"ddns/service"
	"log"
	"strconv"
	"strings"
	"time"
)

func main() {

	for {

		// Buscar o IP publico usando API https://api64.ipify.org
		publicIP, err := service.GetPublicIP()
		if err != nil {
			log.Println("Erro ao processar o método GetPublicIP: ", err)
			continue
		}

		log.Println("IP Público: ", publicIP)

		// Buscar a chave de DNS
		dnsKey, err := service.GetDSNKey()
		if err != nil {
			log.Println("Erro ao processar o método GetDSNKey: ", err)
			continue
		}

		config := service.Config()
		sublist := strings.Split(config.SubDomains, ",")

		// Percorre a lista de subdominios no arquivo env e valida se o subdominio existe na cloudflare - Objetivo é atualizar o IP somente desses dominios.
		for _, sub := range sublist {
			for _, result := range dnsKey.Result {
				if strings.TrimSpace(sub) == result.Name {

					//Atualizar o IP na cloudflare
					err = service.PutPublicIPCloudflare(publicIP, result.ID, strings.TrimSpace(sub))
					if err != nil {
						log.Println("Erro ao atualizar o IP publico na cloudflare: ", err)
						continue
					}
					log.Println("Subdominio: ", result.Name)
				}
			}
		}

		log.Println("IP Público atualizado com sucesso")

		timeConfig, err := strconv.Atoi(config.Time)
		if err != nil {
			log.Println("Erro ao converter o tempo de string para inteiro: ", err)
			continue
		}

		time.Sleep(time.Duration(timeConfig) * time.Second)
	}
}
