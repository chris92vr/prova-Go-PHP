var axios = require("axios").default;

var options = {
  method: 'GET',
  url: 'https://instagram29.p.rapidapi.com/search/avtistkid',
  headers: {
    'x-rapidapi-host': 'instagram29.p.rapidapi.com',
    'x-rapidapi-key': '251669fb4bmshce01bc256be2467p164627jsnefa56c023a18'
  }
};

axios.request(options).then(function (response) {
	console.log(response.data);
}).catch(function (error) {
	console.error(error);
});