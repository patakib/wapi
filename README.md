# Historical Weather API

## Overview
This is a REST API written in Go (Golang). 
The underlying database is PostgreSQL.  
Daily weather data is available for 3 cities: <b>Vienna, Sopron, Budapest</b> - from <b>1980-01-01 to 2023-09-15</b>.

There are two ways to interact with the API: with API-key and without.  
In a real-life scenario, there could be two roles: a <b>basic user and a data analyst</b>.
If you do not use an API key, limited information is available. By using API-key, all the data is available for data analysts.

The architecture of the application is the following:
[](wapi.svg)

## Quickstart

### Dependencies
Docker and Docker Compose. You need to install both on your computer.  

### Environment Variables
Create a `.env.secret` file with the following environment variables:  
```
POSTGRES_USER={YOUR_PG_USER}
POSTGRES_PASS={YOUR_PG_PASS}
POSTGRES_DB={YOUR_PG_DB}
POSTGRES_HOST=localhost
POSTGRES_PORT={YOUR_PG_PORT}
API_KEY={SUPERSECRET_API_KEY}
API_PORT={YOUR_API_PORT}
``` 

Afterwards go into the project folder and start the application and initialize the database:  
`docker compose --env-file .env.secret up -d`  

Make sure both services are available (database and application) and running on the defined ports:  
`docker ps`  

### Endpoints
There are two endpoints for normal users and data analysts:  

1. Limited data for normal users:   
`localhost:{8080}/api/city={CITY[Sopron,Vienna or Budapest]}&date={YYYY-MM-DD}`  
For example:  
`http://localhost:8080/api/city=Sopron&date=2021-01-03`   
Response:  
```
{  
  "city": "Sopron",  
  "date": "2021-01-03T00:00:00Z",  
  "temperature_2m_mean": 3.9,  
  "precipitation_sum": 1.4  
}
```


2. Extended data for data analysts:  
`localhost:{API_PORT}/api/city={CITY[Sopron,Vienna or Budapest]}&date={YYYY-MM-DD}`  
API key sent as "api-key" value in header (or in frontend it can be provided in the form).
For example:  
Response:  
```
{  
  "city": "Sopron",  
  "date": "2021-01-03T00:00:00Z",  
  "temperature_2m_max": 6.4,  
  "temperature_2m_min": 1.6,  
  "temperature_2m_mean": 3.9,  
  "sunrise": "2021-01-03T08:42:00Z",  
  "sunset": "2021-01-03T17:14:00Z",  
  "precipitation_sum": 1.4,  
  "rain_sum": 1.4,  
  "snowfall_sum": 0,  
  "precipitation_hours": 7,  
  "windspeed_10m_max": 14.8,  
  "winddirection_10m_dominant": 156  
}
```


For the full data, you need to attach your <b>{API_KEY}</b> as a value to the key: <b>api-key</b> into the HTTP request header.

### Stop Services and Delete database
`docker compose --env-file .env.secret down --rmi all -v`
Delete `wapi_pgdata` folder (this is the persistent volume of the Postgres container).

### Credits
Credits to open-meteo.com Historical Weather API for the data.
