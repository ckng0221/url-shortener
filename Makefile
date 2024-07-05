run_server_dev:
	CompileDaemon -command="./url-shortener"

create_url:
	curl --location 'localhost:8000/url-shortener' \
		--header 'Content-Type: application/json' \
		--data '{ "url": "www.google.com" }'