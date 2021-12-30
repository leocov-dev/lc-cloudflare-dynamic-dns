FROM golang:1.17-alpine AS builder

WORKDIR /go/src/lc-cloudflare-dynamic-dns
COPY . .
RUN GO111MODULE=on go build -o /go/bin/lc-cf-dns

FROM builder

WORKDIR /
COPY --from=builder /go/bin/lc-cf-dns .

CMD ["lc-cf-dns"]
#ENTRYPOINT ["tail", "-f", "/dev/null"]