.PHONY: install-security
install-security:
	go install github.com/securego/gosec/v2/cmd/gosec@v2.12.0

.PHONY: install-lint
install-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2

.PHONY: lint
lint:
	golangci-lint run -c ./.golangci.yaml

.PHONY: security
security:
	gosec ./...

.PHONY: sdk
sdk:
	cd .. && git clone "https://github.com/aws/aws-sdk-go-v2.git" -b main --depth 1

.PHONY: genapiaws
genapiaws:
	cd genapi && go build && ./genapi "./docs/aws" "../dumper" "../../aws-sdk-go-v2" && go mod tidy
