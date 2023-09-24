# Historical Weather API

## Overview
This is a REST API written in Go (Golang). 
The underlying database is PostgreSQL.  
Daily weather data is available for 3 cities: <b>Vienna, Sopron, Budapest</b> - from <b>1980-01-01 to 2023-09-15</b>.

There are two ways to interact with the API: with API-key and without.  
In a real-life scenarion, there could be two roles: a basic user and a data analyst.
If you do not use an API key, limited information is available. By using API-key, all the data is available for data analysts.

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
`localhost:{API_PORT}/api/daily/{CITY[Sopron,Vienna or Budapest]}/{YYYY-MM-DD}`  

For example:  
`http://localhost:3000/api/daily/Sopron/2021-01-03`  

`localhost:{API_PORT}/api/daily/full/{CITY[Sopron,Vienna or Budapest]}/{YYYY-MM-DD}`  

For example:  
`http://localhost:3000/api/daily/full/Sopron/2021-01-03`  

For the latter you need to add your <b>{API_KEY}</b> as a value to the key: <b>api-key</b> to the HTTP request header.
