run_server_dev:
	CompileDaemon -command="./url-shortener"

create_url:
	curl --location 'localhost:8000/url-shortener' \
		--header 'Content-Type: application/json' \
		--data '{ "url": "https://www.google.com" }'

test_shorten_url:
	curl --location --request GET 'localhost:8000/urls/1'