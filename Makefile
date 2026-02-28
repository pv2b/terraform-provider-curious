default: install

.PHONY: build
build:
	go build -o terraform-provider-curious

.PHONY: install
install: build
	mkdir -p ~/.terraform.d/plugins/registry.terraform.io/pv2b/curious/0.1.0/$$(go env GOOS)_$$(go env GOARCH)
	cp terraform-provider-curious ~/.terraform.d/plugins/registry.terraform.io/pv2b/curious/0.1.0/$$(go env GOOS)_$$(go env GOARCH)/

.PHONY: test
test:
	go test -v ./...

.PHONY: testacc
testacc:
	TF_ACC=1 go test ./... -v -timeout 120m

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: clean
clean:
	rm -f terraform-provider-curious
	rm -rf dist/

.PHONY: docs
docs:
	tfplugindocs generate

.PHONY: version
version:
	@echo "Current version: $$(git describe --tags --always --dirty)"
