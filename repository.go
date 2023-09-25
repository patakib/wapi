package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type DailyWeather struct {
	City                     string    `json:"city"`
	Date                     string    `json:"date"`
	Temp2mMax                float32   `json:"temperature_2m_max"`
	Temp2mMin                float32   `json:"temperature_2m_min"`
	Temp2mMean               float32   `json:"temperature_2m_mean"`
	Sunrise                  time.Time `json:"sunrise"`
	Sunset                   time.Time `json:"sunset"`
	PrecipitationSum         float32   `json:"precipitation_sum"`
	RainSum                  float32   `json:"rain_sum"`
	SnowSum                  float32   `json:"snowfall_sum"`
	PrecipitationHours       int       `json:"precipitation_hours"`
	Windspeed10mMax          float32   `json:"windspeed_10m_max"`
	Winddirection10mDominant int       `json:"winddirection_10m_dominant"`
}

type DailyWeatherReduced struct {
	City             string  `json:"city"`
	Date             string  `json:"date"`
	Temp2mMean       float32 `json:"temperature_2m_mean"`
	PrecipitationSum float32 `json:"precipitation_sum"`
}

type Repository interface {
	GetDailyWeatherWithAuth(city, day string) (*DailyWeather, error)
	GetDailyWeather(city, day string) (*DailyWeatherReduced, error)
}

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(port, host, database, user, pass string) (*PostgresRepository, error) {
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, database)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresRepository{
		db: db,
	}, nil
}

func (pr *PostgresRepository) GetDailyWeather(city, day string) (*DailyWeatherReduced, error) {
	rows, err := pr.db.Query("SELECT city, date, temperature_2m_mean, precipitation_sum_mm FROM weather_daily WHERE city LIKE $1 AND date=$2", city, day)
	if err != nil {
		return nil, err
	}
	dailyWeatherList := []*DailyWeatherReduced{}
	for rows.Next() {
		dailyWeatherItem, err := scanToDailyWeatherReduced(rows)
		if err != nil {
			return nil, err
		}
		dailyWeatherList = append(dailyWeatherList, dailyWeatherItem)
	}
	if len(dailyWeatherList) == 0 {
		return nil, fmt.Errorf(
			"Weather data not found for these parameters, city: %s, date: %s. Weather data is available for Vienna, Budapest and Sopron from 1980-01-01 to 2023-09-15.",
			day,
			city,
		)
	}
	return dailyWeatherList[0], nil
}

func (pr *PostgresRepository) GetDailyWeatherWithAuth(city, day string) (*DailyWeather, error) {
	rows, err := pr.db.Query("SELECT * FROM weather_daily WHERE city LIKE $1 AND date=$2", city, day)
	if err != nil {
		return nil, err
	}
	dailyWeatherList := []*DailyWeather{}
	for rows.Next() {
		dailyWeatherItem, err := scanToDailyWeather(rows)
		if err != nil {
			return nil, err
		}
		dailyWeatherList = append(dailyWeatherList, dailyWeatherItem)
	}
	if len(dailyWeatherList) == 0 {
		return nil, fmt.Errorf(
			"Weather data not found for these parameters, city: %s, date: %s. Weather data is available for Vienna, Budapest and Sopron from 1980-01-01 to 2023-09-15.",
			day,
			city,
		)
	}
	return dailyWeatherList[0], nil
}

func scanToDailyWeatherReduced(rows *sql.Rows) (*DailyWeatherReduced, error) {
	dailyWeatherItem := new(DailyWeatherReduced)
	err := rows.Scan(
		&dailyWeatherItem.City,
		&dailyWeatherItem.Date,
		&dailyWeatherItem.Temp2mMean,
		&dailyWeatherItem.PrecipitationSum)

	return dailyWeatherItem, err
}

func scanToDailyWeather(rows *sql.Rows) (*DailyWeather, error) {
	dailyWeatherItem := new(DailyWeather)
	err := rows.Scan(
		&dailyWeatherItem.City,
		&dailyWeatherItem.Date,
		&dailyWeatherItem.Temp2mMax,
		&dailyWeatherItem.Temp2mMin,
		&dailyWeatherItem.Temp2mMean,
		&dailyWeatherItem.Sunrise,
		&dailyWeatherItem.Sunset,
		&dailyWeatherItem.PrecipitationSum,
		&dailyWeatherItem.RainSum,
		&dailyWeatherItem.SnowSum,
		&dailyWeatherItem.PrecipitationHours,
		&dailyWeatherItem.Windspeed10mMax,
		&dailyWeatherItem.Winddirection10mDominant)

	return dailyWeatherItem, err
}
