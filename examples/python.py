import os
import requests

api_key = os.environ.get("JSONCARGO_API_KEY")
if not api_key:
    raise EnvironmentError("JSONCARGO_API_KEY environment variable is not set")

tracking_number = "MSCU1234567"  # replace with a real container number

url = f"http://api.jsoncargo.com/api/v1/containers/{tracking_number}"

headers = {
    "x-api-key": api_key
}

# If the container prefix is shared across carriers, add shipping_line:
# params = {"shipping_line": "MSC"}
params = {}

response = requests.get(url, headers=headers, params=params)

if response.status_code == 200:
    data = response.json()["data"]
    print(f"Container:  {data['container_id']}")
    print(f"Status:     {data['container_status']}")
    print(f"From:       {data['shipped_from']}")
    print(f"To:         {data['shipped_to']}")
    print(f"ETA:        {data['eta_final_destination']}")
else:
    print(f"Error {response.status_code}: {response.text}")
