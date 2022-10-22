FROM golang:1.17-alpine AS builder

WORKDIR /code
COPY . .
RUN go build -o ./bin/cloudflare-dynamic-dns

FROM builder

WORKDIR /app
COPY --from=builder /code/bin/cloudflare-dynamic-dns .

CMD ["cloudflare-dynamic-dns"]