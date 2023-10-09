let dataMap = {
    "date": "Date",
    "city": "City",
    "temperature_2m_max": "Max. Temperature (2m)",
    "temperature_2m_min": "Min. Temperature (2m)",
    "temperature_2m_mean": "Mean Temperature (2m)",
    "sunrise": "Sunrise",
    "sunset": "Sunset",
    "precipitation_sum": "Precipitation sum",
    "rain_sum": "Rain",
    "snowfall_sum": "Snowfall",
    "precipitation_hours": "Precipitation Hours",
    "windspeed_10m_max": "Max. Windspeed (10m)",
    "winddirection_10m_dominant": "Dominant Wind Direction (10m)"
}

async function getWeather() {
    let baseUrl = "/api/";
    let city = document.getElementById("city").value;
    let date = document.getElementById("date").value;
    let apiKey = document.getElementById("apikey").value;
    suffix = `city=${city}&date=${date}`
    url = baseUrl.concat(suffix)

    const response = await fetch(url, {
        method: "GET",
        headers: {
            'api-key': apiKey
        }
    });
    const weather = await response.json();
    let tableBody = document.getElementById("results");
    tableBody.innerHTML = "";
    for (let key in weather) {
        let row = document.createElement("tr")
        let field = document.createElement("td")
        let value = document.createElement("td")
        field.innerText = dataMap[key]
        value.innerText = weather[key] 
        row.appendChild(field)
        row.appendChild(value)
        tableBody.appendChild(row)
    }
};

async function forecastWeather() {
    let baseUrl = "/api/forecast/";
    let city = document.getElementById("city").value;
    let date = document.getElementById("date").value;
    let apiKey = document.getElementById("apikey").value;
    suffix = `city=${city}&date=${date}`
    url = baseUrl.concat(suffix)

    const response = await fetch(url, {
        method: "GET",
        headers: {
            'api-key': apiKey
        }
    });
    const weather = await response.json();
    let tableBody = document.getElementById("results");
    tableBody.innerHTML = "";
    for (let key in weather) {
        let row = document.createElement("tr")
        let field = document.createElement("td")
        let value = document.createElement("td")
        field.innerText = dataMap[key]
        value.innerText = weather[key] 
        row.appendChild(field)
        row.appendChild(value)
        tableBody.appendChild(row)
    }
};





