GOFMT_FILES?=$$(find . -type f -name '*.go')

default: dev

dev: clean tidy fmt
	@go build -race -o "bin/cloudflare-dynamic-dns" .

docker-image:
	@docker build -t cloudflare-dynamic-dns .

# bin generates release zip packages in ./dist
release: clean tidy
	@sh -c "scripts/release.sh"

clean:
	@rm -rf "bin"
	@rm -rf "dist"

fmt:
	@gofmt -w $(GOFMT_FILES)

fmtcheck:
	@bash -c "scripts/gofmtcheck.sh"

tidy:
	@go mod tidy

.NOTPARALLEL:

.PHONY: bin default dev fmtcheck fmt tidy