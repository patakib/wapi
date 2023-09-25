package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetDailyWeather(t *testing.T) {
}

func TestGetDailyWeatherWithAuth(t *testing.T) {

}

func TestScanToDailyWeather(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Println("Error creating a mock db.")
	}
	defer db.Close()
	rows := mock.NewRows(
		[]string{
			"city",
			"date",
			"temperature_2m_max",
			"temperature_2m_min",
			"temperature_2m_mean",
			"sunrise",
			"sunset",
			"precipitation_sum",
			"rain_sum",
			"snow_sum",
			"precipitation_hours",
			"windspeed_10m_max",
			"winddirection_10m_dominant",
		},
	).AddRow(
		"Sopron",
		"2020-01-01",
		0.5,
		-6.5,
		-2.1,
		time.Now(),
		time.Now(),
		3.4,
		3.3,
		0.1,
		4,
		21.5,
		130,
	)
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	rs, _ := db.Query("SELECT")
	testDailyWeatherList := []*DailyWeather{}
	for rs.Next() {
		testDailyWeather, err := scanToDailyWeather(rs)
		if err != nil {
			fmt.Printf("%s, Error scanning to daily weather.", err)
		}
		testDailyWeatherList = append(testDailyWeatherList, testDailyWeather)
	}
	if testDailyWeatherList[0].City != "Sopron" {
		t.Errorf("Row incorrectly parsed into struct!")
	}
}

func TestScanToDailyWeatherReduced(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Println("Error creating a mock db.")
	}
	defer db.Close()
	rows := mock.NewRows(
		[]string{
			"city",
			"date",
			"temperature_2m_mean",
			"precipitation_sum",
		},
	).AddRow(
		"Sopron",
		"2020-01-01",
		-2.1,
		3.4,
	)
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	rs, _ := db.Query("SELECT")
	testDailyWeatherReducedList := []*DailyWeatherReduced{}
	for rs.Next() {
		testDailyWeatherReduced, err := scanToDailyWeatherReduced(rs)
		if err != nil {
			fmt.Printf("%s, Error scanning to daily weather.", err)
		}
		testDailyWeatherReducedList = append(testDailyWeatherReducedList, testDailyWeatherReduced)
	}
	if testDailyWeatherReducedList[0].City != "Sopron" {
		t.Errorf("Row incorrectly parsed into struct!")
	}
}
