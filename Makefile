PKG_NAME=microscaler-flux

default: build

dependencies:
	GO111MODULE=on go install github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

build:
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -o ./bin/terraform-provider-${PKG_NAME} -a -ldflags '-extldflags "-static"' main.go

install: build
	mkdir -p ~/.terraform.d/plugins/registry.terraform.io/microscaler/terraform-provider-${PKG_NAME}/0.1.0/$${GOOS}_$${GOARCH}
	cp ./bin/flux ~/.terraform.d/plugins/registry.terraform.io/microscaler/terraform-provider-${PKG_NAME}/0.1.0/$${GOOS}_$${GOARCH}/terraform-provider-${PKG_NAME}

init:
	terraform init

apply: init
	terraform apply

fmt:
	go fmt ./...

tidy:
	rm -f go.sum; go mod tidy -compat=1.18

test: tidy fmt
	go test ./...

.PHONY: install build
