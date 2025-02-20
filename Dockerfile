# Etapa 1: Usar a imagem base do Golang
FROM golang:1.23.3 AS builder

# Definir o diretório de trabalho no container
WORKDIR /ddns_ip_cloudflare

# Copiar go.mod e go.sum
COPY go.mod go.sum ./

# Baixar as dependências do Go
RUN go mod download

# Copiar todo o código para dentro do container (a partir da raiz do projeto)
COPY . .

# Compilar o binário da aplicação
RUN go build -o ddns . 

# Etapa 2: Criar uma imagem final com a aplicação compilada
FROM golang:1.23.3

# Definir o diretório de trabalho no container
WORKDIR /ddns_ip_cloudflare

# Copiar o binário compilado da imagem builder
COPY --from=builder /ddns_ip_cloudflare/ddns /ddns_ip_cloudflare/

# Expor a porta onde a aplicação vai rodar
#EXPOSE 3000

# Comando para rodar a aplicação
CMD ["./ddns"]

