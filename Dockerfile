FROM golang:1.24.6

WORKDIR /src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 8182

RUN go build -o main cmd/main.go

CMD ["./main"]