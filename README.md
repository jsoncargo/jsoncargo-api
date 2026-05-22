# JSONCargo API

OpenAPI spec and code examples for the [JSONCargo](https://jsoncargo.com) container and vessel tracking API.

For the full documentation, go to [jsoncargo.com/documentation-api](https://jsoncargo.com/documentation-api/).

## What is in this repo

- `openapi.yaml` — complete OpenAPI 3.0 spec covering all 10 endpoints, request parameters, and response schemas
- `examples/` — one working code sample per language for a basic container lookup

## Quickstart

Replace `YOUR_CONTAINER_NUMBER` with a real container number (e.g. `MSCU1234567`). Store your API key in the `JSONCARGO_API_KEY` environment variable.

```bash
export JSONCARGO_API_KEY=your_api_key_here

curl -X GET "http://api.jsoncargo.com/api/v1/containers/MSCU1234567" \
  -H "x-api-key: $JSONCARGO_API_KEY"
```

If the container prefix is shared across multiple carriers, add the `shipping_line` param:

```bash
curl -X GET "http://api.jsoncargo.com/api/v1/containers/CCLU1234567?shipping_line=COSCO" \
  -H "x-api-key: $JSONCARGO_API_KEY"
```

## Supported shipping lines

`MAERSK` `HAPAG_LLOYD` `HMM` `ONE` `EVERGREEN` `MSC` `CMA_CGM` `COSCO` `ZIM` `YANG_MING` `PIL`

## Using the OpenAPI spec

Import `openapi.yaml` directly into Postman: click **Import** at the top of the Collections sidebar (or **File > Import**), select the file, and confirm. Postman will generate a request for each endpoint with all parameters pre-filled.

You can also use it with any OpenAPI-compatible tool (Insomnia, Swagger UI, code generators, etc.).

## Support

- support@jsoncargo.com
- [jsoncargo.com/documentation-api](https://jsoncargo.com/documentation-api/)
