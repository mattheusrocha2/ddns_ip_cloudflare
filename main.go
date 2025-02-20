package main

import (
	"log"
	"net"
)

func main() {

	// Verificar o IP da minha maquina e consultar o seu IP publico.
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Println("Erro para obter o endereço IP Local. ", err)
		return
	}

	for _, addr := range addrs {
		// if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
		// 	fmt.Println("IP Local: ", ipnet.IP.String())
		// 	fmt.Println("IP Público: ", getPublicIP())
		// }

		ipnet, _ := addr.(*net.IPNet)
		ipnet.IP.IsLoopback()
		

	}

	// Inicializa o método responsável por enviar a solicitação de atualização do IP público para o serviço Cloudflare.

}
