deps:
	go mod vendor

build:
	go build -o bin/webhook-plex cmd/main.go

clean:
	rm -r bin

test:
	go test -timeout 30s ./plugins/inputs/webhooks/plex

run:
	./bin/webhook-plex --config plugin.conf
