.DEFAULT_GOAL := build

build:
	go build -o bin/jamf_webhooks cmd/main.go

deps:
	go mod vendor

clean:
	rm -r bin

test:
	go test -timeout 30s -count=1 ./plugins/inputs/webhooks/jamf

run:
	./bin/jamf_webhooks --config plugin.conf
