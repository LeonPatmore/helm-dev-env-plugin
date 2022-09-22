setup:
	export AWS_PROFILE=nexmo-dev
	helm repo add olympus-service s3://nexmo-core-helm/olympus/olympus-service/charts
	go get .

uninstall:
	helm plugin uninstall olympus-dev || true

install: uninstall
	echo "Ensure you are running as admin!"
	helm plugin install .

test-helm:
	helm olympus-dev

run-go:
	go run main.go

build:
	go build

test:
	go test
