.DEFAULT_GOAL := build

build:
	go build -o bin/webhook-jamf cmd/main.go

deps:
	go mod vendor

clean:
	rm -r bin

test:
	go test -timeout 30s -count=1 ./plugins/inputs/webhooks/jamf

run:
	./bin/webhook-jamf --config plugin.conf
