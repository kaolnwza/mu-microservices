FROM golang:1.17.4 AS builder

WORKDIR /go/app

COPY . .

# COPY go.mod .

# COPY go.sum .

RUN go mod tidy -go=1.16 && go mod tidy -go=1.17

RUN cd cmd && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build main.go grpc.server.go
# RUN cd cmd && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /build

FROM alpine:3.13

WORKDIR /app

# COPY --from=builder /go/app/credentials/ssl/ca.crt ./credentials/ssl/ca.crt
# COPY --from=builder /go/app/credentials/ssl/server.crt ./credentials/ssl/server.crt
# COPY --from=builder /go/app/credentials/ssl/server.pem ./credentials/ssl/server.pem

COPY --from=builder /go/app/cmd/main main

COPY --from=builder /go/app/credentials/long-sonar-371616-80145143437f.json ./credentials/long-sonar-371616-80145143437f.json

#old:new
# EXPOSE 8001 50052

# CMD ["./go/app/apisrv/main"]

CMD ["./main"]
