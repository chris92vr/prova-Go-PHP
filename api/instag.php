<?php

$curl = curl_init();

curl_setopt_array($curl, [
	CURLOPT_URL => "https://instagram29.p.rapidapi.com/search/avtistkid",
	CURLOPT_RETURNTRANSFER => true,
	CURLOPT_FOLLOWLOCATION => true,
	CURLOPT_ENCODING => "",
	CURLOPT_MAXREDIRS => 10,
	CURLOPT_TIMEOUT => 30,
	CURLOPT_HTTP_VERSION => CURL_HTTP_VERSION_1_1,
	CURLOPT_CUSTOMREQUEST => "GET",
	CURLOPT_HTTPHEADER => [
		"x-rapidapi-host: instagram29.p.rapidapi.com",
		"x-rapidapi-key: 251669fb4bmshce01bc256be2467p164627jsnefa56c023a18"
	],
]);

$response = curl_exec($curl);
$err = curl_error($curl);

curl_close($curl);

if ($err) {
	echo "cURL Error #:" . $err;
} else {
	echo $response;
}