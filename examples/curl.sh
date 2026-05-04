#!/bin/bash

# Set your API key as an environment variable:
#   export JSONCARGO_API_KEY=your_api_key_here

if [ -z "$JSONCARGO_API_KEY" ]; then
  echo "Error: JSONCARGO_API_KEY environment variable is not set" >&2
  exit 1
fi

TRACKING_NUMBER="MSCU1234567" # replace with a real container number

# If the container prefix is shared across carriers, add ?shipping_line=MSC
curl -s -X GET \
  "http://api.jsoncargo.com/api/v1/containers/${TRACKING_NUMBER}" \
  -H "x-api-key: ${JSONCARGO_API_KEY}" | python3 -m json.tool
