.PHONY: test test_v test_cpf test_cov

test:
	go test ./...

test_v:
	go test -v ./...

test_cov:
	go test ./... -cover
