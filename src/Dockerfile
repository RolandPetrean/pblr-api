FROM golang:1.19-alpine

WORKDIR /src

COPY . .

RUN go mod download

CMD go run .