<?php

$client = new http\Client;
$request = new http\Client\Request;

$request->setRequestUrl('https://instagram-data1.p.rapidapi.com/location/feed');
$request->setRequestMethod('GET');
$request->setQuery(new http\QueryString([
	'location_id' => '213385402',
	'end_cursor' => '2312589909032003834'
]));

$request->setHeaders([
	'x-rapidapi-host' => 'instagram-data1.p.rapidapi.com',
	'x-rapidapi-key' => '251669fb4bmshce01bc256be2467p164627jsnefa56c023a18'
]);

$client->enqueue($request)->send();
$response = $client->getResponse();

echo $response->getBody();