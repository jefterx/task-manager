FROM golang:1.23-alpine

# Definir o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copiar o arquivo go.mod e go.sum para o diretório de trabalho
COPY go.mod go.sum ./

# Baixar as dependências
RUN go mod download

# Copiar o código-fonte para o diretório de trabalho
COPY . .

# Compilar a aplicação
RUN go build -o main .

# Expor a porta em que a aplicação será executada
EXPOSE 8000

# Comando para executar a aplicação
CMD ["./main"]
