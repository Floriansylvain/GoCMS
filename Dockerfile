FROM golang:1.20.5-alpine

WORKDIR /app

COPY go.mod go.sum ./
COPY adapters ./adapters
COPY api ./api
COPY domain ./domain
COPY main ./main
COPY useCases ./useCases

RUN go mod download
RUN go build -o startServer ./main

CMD ["./startServer"]
