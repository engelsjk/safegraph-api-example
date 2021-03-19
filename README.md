# safegraph-api-example

A quick test of the [Safegraph Places API](https://docs.safegraph.com/reference).

## Run

First, you'll need an .env file with a SAFEGRAPH_API_KEY variable.

Then, to run a sample request...

```bash
go run example.go

```

Compare that output with the same request using curl...

```bash
curl 'https://api.safegraph.com/v1/graphql' \
-H 'apikey: {SAFEGRAPH_API_KEY}' \
-H 'content-type: application/json' \
--data-raw '{
	"query": "query($placekeys: [Placekey!]) {\n\t\tplaces(placekeys: $placekeys) {\n\t\t\tplacekey\n\t\t\tsafegraph_core {\n\t\t\t\tlocation_name\n\t\t\t\ttop_category\n\t\t\t\tstreet_address\n\t\t\t\tcity\n\t\t\t\tregion\n\t\t\t\tlatitude\n\t\t\t\tlongitude\n\t\t\t}\n\t\t}\n\t}",
	"variables": "{\n\t\t\"placekeys\": [\n\t\t\t\"222-222@5qw-shj-7qz\",\n\t\t\t\"222-222@5s6-pyc-7qz\",\n\t\t\t\"zzw-222@5vg-3tv-7qz\",\n\t\t\t\"22s-223@63j-4f8-7qz\",\n\t\t\t\"228-222@5sb-cyn-7qz\"\n\t\t]\n\t}"
}' 
```