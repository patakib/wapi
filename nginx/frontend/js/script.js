async function getWeather() {
    let baseUrl = "/api/daily/";
    let city = document.getElementById("city").value;
    let date = document.getElementById("date").value;
    let apiKey = document.getElementById("apikey").value;
    if (apiKey != null || apiKey != "") {
        suffix = `full/${city}/${date}`
        url = baseUrl.concat(suffix)
    } else {
        suffix = `${city}/${date}`
        url = baseUrl.concat(suffix)
    }

    const response = await fetch(url, {
        method: "GET",
        headers: {
            'api-key': apiKey,
            'accept': 'application/json',
          },
    });

    const weather = await response.json();
    console.log(weather);
    }

getWeather();