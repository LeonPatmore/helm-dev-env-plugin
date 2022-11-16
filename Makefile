setup:
	go get .
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	helm repo add bitnami https://charts.bitnami.com/bitnami
	minikube start

uninstall:
	helm plugin uninstall dev || true

install: uninstall
	echo "Ensure you are running as admin!"
	helm plugin install .

test-helm:
	export AWS_REGION="eu-west-1" && helm dev --devname leon -t leon-dev

run-go:
	go run .

build:
	go build

test:
	go test -v -coverprofile cover.out

coverage:
	go tool cover -html cover.out

lint:
	golangci-lint run --timeout=3m

format:
	gofmt -s -w .
