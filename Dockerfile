FROM golang:1.24.6

WORKDIR /src/app

RUN mkdir -p /data/db && chown -R mongodb:mongodb /data/db && chmod -R 770 /data/db

VOLUME /data/db

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 8182

RUN go build -o main cmd/main.go

CMD ["./main"]