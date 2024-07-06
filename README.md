# URL Shortener

A URL Shortener API, that creates shorten URL based on user input. The shorten URL is based on `base62` character using `0-9`, `a-z`, `A-Z`.

## How it works

The URL Shortener uses distributed 64-bit unique ID generator (based on Twitter's Snowflake ID) to generate an ID for the URL record. The ID is converted into integer, and then converted to its base62 representative, and stored as the shorten URL path.

Eg.

- ID: `7215288079162728448`
- Base62: `8B05WPRMQ7e`
- Shorten URL Path: `{BASE_URL}/urls/8B05WPRMQ7e`

Accessing to the shorten URL path will redirect the client to the orignal URL, also updating the usage count and accessed datetime.

## Get started

Rename the `.env.example` to `.env`, and update the values accordingly.

```bash
# Install dependencies
go get .
go get github.com/githubnemo/CompileDaemon

# Run server (dev mode)
make run_server_dev
# Run server (without watch)
go run .

# Test APIS
make create_url

# Get shorten URL record, eg.
make get_shorten_url ID=7215288128387080192

# Use use_shorten_url, eg.
make use_shorten_url shortenUrl=8B05XHB5cEo


# Run server (prod mode)
make build
make run_server_prod
```
