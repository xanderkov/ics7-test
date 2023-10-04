FROM golang:1.20-alpine
WORKDIR /build

COPY . .

RUN go mod tidy
#RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/server/main.go

CMD ["go test ./..."]