FROM golang:1.19-alpine3.17 AS builder

ENV CGO_ENABLED=0

WORKDIR /code
COPY . .
RUN go build \
      -o ./bin/cloudflare-dynamic-dns \
      --ldflags="-s -w -X 'lc-cloudflare-dynamic-dns/config.Version=${RELEASE_TAG}'"

FROM alpine:3.17

WORKDIR /app
COPY --from=builder /code/bin/cloudflare-dynamic-dns .

ENTRYPOINT ["./cloudflare-dynamic-dns", "update"]