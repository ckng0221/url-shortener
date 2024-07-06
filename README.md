# URL Shortener

A URL Shortener API web API, that creates shorten URL based on user's input. The shorten URL is created using `base62` characters, ranging from `0-9`, `a-z`, `A-Z` (total of 62 characters).

## How it works

The application has a distributed 64-bit unique and sortable ID generator (based on Twitter's [Snowflake ID](https://en.wikipedia.org/wiki/Snowflake_ID)) that is used to generate a unique ID for each URL record. The structure of ID is as follows:

| Component | 0   | timestamp | datacenterID | machineID | sequenceNumber |
| --------- | --- | --------- | ------------ | --------- | -------------- |
| Bit(s)    | 1   | 41        | 5            | 5         | 12             |

- `0`: Placeholder, could indicate the sign, remain for future use.
- `timestamp`: Unix timestamp in milliseconds.
- `dataCenterID`: Data Center ID. Can have maximum of 32 data center IDs.
- `machineId`: Machine ID. Can have maximum of 32 machine (nodes) in each data center.
- `sequenceNumber`: Sequence number. Can have maximum of 4096 sequence number per millisecond.

Example generated ID:

- Binary: `0110010000100001110110100100011110011110100110101110000000000000`
- Decimal: `7215288079162728448`

Upon user's request, a unique ID is generated. The ID is then converted to its base62 representative, and stored as the shorten URL path.

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
