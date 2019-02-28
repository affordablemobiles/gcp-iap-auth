FROM golang:1.12.0-alpine3.9 AS builder
RUN apk add --no-cache git
WORKDIR /build
COPY . .
RUN go build .

FROM alpine:3.9
RUN apk add --no-cache ca-certificates \
    && addgroup -S gcp-iap-auth \
    && adduser -S gcp-iap-auth -G gcp-iap-auth \
    && mkdir /app \
    && chown gcp-iap-auth:gcp-iap-auth /app
USER gcp-iap-auth
WORKDIR /app
COPY --from=builder --chown=gcp-iap-auth /build/gcp-iap-auth .

CMD ["/app/gcp-iap-auth"]
