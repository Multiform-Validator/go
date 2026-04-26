GOCACHE ?= /tmp/go-build
STATICCHECK_CACHE ?= /tmp/staticcheck

export GOCACHE
export STATICCHECK_CACHE

.PHONY: fmt fmt_check vet staticcheck security lint check test test_v test_cov coverage coverage_html clean_cov

test:
	go test ./...

test_v:
	go test -v ./...

test_cov:
	go test ./... -coverprofile=cover.out -cover
	go tool cover -func=cover.out

coverage: test_cov coverage_html

coverage_html:
	go tool cover -html=cover.out -o coverage.html

fmt:
	gofmt -w .

fmt_check:
	@test -z "$$(gofmt -l .)" || (gofmt -l . && exit 1)

vet:
	go vet ./...

staticcheck:
	@if command -v staticcheck >/dev/null 2>&1; then \
		staticcheck ./...; \
	else \
		echo "staticcheck not installed; skipping"; \
	fi

security:
	@if command -v govulncheck >/dev/null 2>&1; then \
		govulncheck ./...; \
	else \
		echo "govulncheck not installed; skipping"; \
	fi
	@if command -v gosec >/dev/null 2>&1; then \
		gosec ./...; \
	else \
		echo "gosec not installed; skipping"; \
	fi

lint: fmt_check vet staticcheck

check: lint security test_cov coverage_html

clean_cov:
	rm -f cover.out coverage.html
