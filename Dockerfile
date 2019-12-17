FROM golang:1.13-alpine as builder

RUN mkdir -p /go/src/github.com/bladedancer/xdsing

WORKDIR /go/src/github.com/bladedancer/xdsing

# Copy necessary files
ADD . .

RUN rm -rf bin
RUN CGO_ENABLED=0 GOOS=linux  GOARCH=amd64 go build -o bin/xdsing ./cmd/xdsing

# Create non-root user
RUN addgroup -S envoy && adduser -S envoy -G envoy
RUN chown -R envoy:envoy /go/src/github.com/bladedancer/xdsing/bin/xdsing
USER envoy

# Base image
FROM scratch

# Copy binary and user from previous build step
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github.com/bladedancer/xdsing/bin/xdsing /root/xdsing
COPY --from=builder /etc/passwd /etc/passwd
USER envoy

ENTRYPOINT ["/root/xdsing"]
