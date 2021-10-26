FROM golang:1.17-alpine

WORKDIR /annonce/

COPY . ./

RUN go mod download

RUN go build -o ./cmd/main ./cmd/main.go 

ENTRYPOINT [ "./cmd/main" ]

ENV PORT=8081