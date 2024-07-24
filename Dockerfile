# Building the Golang Binary
FROM golang:alpine AS builder
LABEL stage=auto-clean
ENV CGO_ENABLED 0
WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN go build -o /go/bin/app ./cmd/main.go

# Final Minimal APP
FROM scratch
COPY --from=builder /go/bin/app /go/bin/app

EXPOSE 8080
ENTRYPOINT ["/go/bin/app"]

CMD ["/go/bin/app"]