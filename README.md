# Build
go build -ldflags "-w -s" -o bin/jamf_webhooks cmd/main.go

# Run
bin/jamf_webhooks --config plugin.conf
