run_server_dev:
	CompileDaemon -command="./url-shortener"

create_url:
	curl --location 'localhost:8000/shorten-urls' \
		--header 'Content-Type: application/json' \
		--data '{ "url": "https://www.google.com" }'

get_shorten_url:
	curl --location --request GET 'localhost:8000/shorten-urls/7215288128387080192'

test_shorten_url:
	curl --location --request GET 'localhost:8000/urls/8B05XHB5cEo'