CREATE TABLE weather_daily (
    city VARCHAR(50),
    date DATE,
    temperature_2m_max DECIMAL,
    temperature_2m_min DECIMAL,
    temperature_2m_mean DECIMAL,
    sunrise TIMESTAMP,
    sunset TIMESTAMP,
    precipitation_sum_mm DECIMAL,
    rain_sum_mm DECIMAL,
    snowfall_sum_cm DECIMAL,
    precipitation_hours INTEGER,
    windspeed_10m_max DECIMAL,
    winddirection_10m_dominant INTEGER
);

COPY weather_daily (city, date, temperature_2m_max, temperature_2m_min, temperature_2m_mean, sunrise, sunset, precipitation_sum_mm, rain_sum_mm, snowfall_sum_cm, precipitation_hours, windspeed_10m_max, winddirection_10m_dominant) FROM '/var/lib/postgresql/csvs/weather_daily.csv' DELIMITER ',' CSV HEADER;