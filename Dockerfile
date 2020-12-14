# golang alpine 1.15.6
FROM golang:1.15.6-alpine as builder

# SSL for HTTPS calls.
RUN apk update && apk add --no-cache ca-certificates tzdata && update-ca-certificates

WORKDIR $GOPATH/src/github.com/javiertlopez/golang-bootcamp-2020
COPY . .

RUN go get github.com/javiertlopez/golang-bootcamp-2020
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/main .

# Small image
FROM scratch

# Import from builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy our static executable.
COPY --from=builder /go/bin/main /go/bin/main

# Port on which the service will be exposed.
EXPOSE 8080

ENTRYPOINT ["/go/bin/main"]