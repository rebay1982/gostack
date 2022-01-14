FROM golang:1.17.5-alpine3.15 as builder
COPY go.mod go.sum /go/src/github.com/rebay1982/gostack/
WORKDIR /go/src/github.com/rebay1982/gostack
RUN go mod download
COPY . /go/src/github.com/rebay1982/gostack
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/gostack github.com/rebay1982/gostack

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/rebay1982/gostack/build/gostack /usr/bin/gostack
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/gostack"]

