run:
	docker-compose -f docker-compose.yml up -d --wait \
    	&& go run main.go
