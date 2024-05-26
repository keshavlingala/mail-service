FROM golang:1.21.6-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -ldflags '-s -w' -o app

EXPOSE 8080

CMD ["./app"]
