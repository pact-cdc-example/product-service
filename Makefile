run:
	docker-compose -f docker-compose.yml up -d --wait \
    	&& go run main.go

test-provider:
	go clean -testcache && go test -tags=provider ./app/... -v