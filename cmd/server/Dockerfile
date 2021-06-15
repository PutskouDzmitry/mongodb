# syntax=docker/dockerfile:1
FROM golang:1.16 as builder
RUN mkdir /app
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -mod=mod -a -installsuffix cgo -o ./cmd/server/cmd.go .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/alexellis/href-counter/app .
CMD ["./app"]  
