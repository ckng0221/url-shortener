# Run server
run_server_dev:
	CompileDaemon -command="./url-shortener"

test_all: test build

test:
	go test ./... 

build:
	go build .

run_server_prod:
	export GIN_MODE=release && \
	./url-shortener

## Run API
create_url:
	curl --location 'localhost:8000/shorten-urls' \
		--header 'Content-Type: application/json' \
		--data '{ "url": "https://www.google.com" }'

# eg. make get_shorten_url ID=7215288128387080192
get_shorten_url:
	curl --location --request GET 'localhost:8000/shorten-urls/$(ID)'

# eg. make use_shorten_url shortenUrl=8B05XHB5cEo
use_shorten_url:
	curl --location --request GET 'localhost:8000/urls/$(shortenUrl)'