<?php

$apiKey = getenv('JSONCARGO_API_KEY');
if (!$apiKey) {
    fwrite(STDERR, "JSONCARGO_API_KEY environment variable is not set\n");
    exit(1);
}

$trackingNumber = 'MSCU1234567'; // replace with a real container number

// If the container prefix is shared across carriers, append ?shipping_line=MSC
$url = 'http://api.jsoncargo.com/api/v1/containers/' . urlencode($trackingNumber);

$curl = curl_init();
curl_setopt_array($curl, [
    CURLOPT_URL            => $url,
    CURLOPT_RETURNTRANSFER => true,
    CURLOPT_HTTPHEADER     => [
        'x-api-key: ' . $apiKey,
    ],
]);

$response = curl_exec($curl);
$statusCode = curl_getinfo($curl, CURLINFO_HTTP_CODE);
curl_close($curl);

if ($statusCode === 200) {
    $data = json_decode($response, true)['data'];
    echo "Container: " . $data['container_id'] . "\n";
    echo "Status:    " . $data['container_status'] . "\n";
    echo "From:      " . $data['shipped_from'] . "\n";
    echo "To:        " . $data['shipped_to'] . "\n";
    echo "ETA:       " . $data['eta_final_destination'] . "\n";
} else {
    fwrite(STDERR, "Error $statusCode: $response\n");
    exit(1);
}
