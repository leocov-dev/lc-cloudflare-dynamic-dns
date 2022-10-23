GOFMT_FILES?=$$(find . -type f -name '*.go')

default: dev

dev: clean tidy fmt
	@go build -race -o "$(CURDIR)/bin/cloudflare-dynamic-dns" .

# bin generates release zip packages in ./dist
release: clean tidy
	@sh -c "$(CURDIR)/scripts/release.sh"

clean:
	@rm -rf "$(CURDIR)/bin"
	@rm -rf "$(CURDIR)/dist"

fmt:
	@gofmt -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

tidy:
	@go mod tidy

.NOTPARALLEL:

.PHONY: bin default dev fmtcheck fmt tidy