const https = require("http"); // API uses http

const apiKey = process.env.JSONCARGO_API_KEY;
if (!apiKey) {
  console.error("JSONCARGO_API_KEY environment variable is not set");
  process.exit(1);
}

const trackingNumber = "MSCU1234567"; // replace with a real container number

// If the container prefix is shared across carriers, add ?shipping_line=MSC
const options = {
  hostname: "api.jsoncargo.com",
  path: `/api/v1/containers/${trackingNumber}`,
  method: "GET",
  headers: {
    "x-api-key": apiKey,
  },
};

const req = https.request(options, (res) => {
  let body = "";
  res.on("data", (chunk) => (body += chunk));
  res.on("end", () => {
    if (res.statusCode === 200) {
      const data = JSON.parse(body).data;
      console.log("Container: ", data.container_id);
      console.log("Status:    ", data.container_status);
      console.log("From:      ", data.shipped_from);
      console.log("To:        ", data.shipped_to);
      console.log("ETA:       ", data.eta_final_destination);
    } else {
      console.error(`Error ${res.statusCode}:`, body);
    }
  });
});

req.on("error", (err) => console.error("Request error:", err.message));
req.end();
