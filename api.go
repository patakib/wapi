package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type customApiFunc func(http.ResponseWriter, *http.Request) error

type ServerError struct {
	Error string
}

func customApiFuncDecorator(caf customApiFunc) http.HandlerFunc {
	// A middleware function to enable logging and error return from ServeHTTP method
	return func(w http.ResponseWriter, r *http.Request) {
		file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
		log.SetOutput(file)
		start := time.Now()
		if err := caf(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ServerError{Error: err.Error()})
		}
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	}
}

type Server struct {
	listenAddress string
	repository    Repository
}

func NewAPIServer(listenAddress string, repository Repository) *Server {
	return &Server{
		listenAddress: listenAddress,
		repository:    repository,
	}
}

func (server *Server) Run() {
	router := mux.NewRouter()
	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/daily/{city}/{day}", customApiFuncDecorator(server.handleGetDailyWeather))
	api.HandleFunc("/daily/full/{city}/{day}", customApiFuncDecorator(server.handleGetDailyWeatherWithAuth))

	log.Println("REST API Server is up and running on port: ", server.listenAddress)

	http.ListenAndServe(server.listenAddress, router)
}

func GetApiKey() string {
	envErr := godotenv.Load(".env.secret")
	if envErr != nil {
		log.Fatal(envErr)
	}
	return os.Getenv("API_KEY")
}

func (server *Server) handleGetDailyWeather(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		vars := mux.Vars(r)
		dailyWeatherList, err := server.repository.GetDailyWeather(vars["city"], vars["day"])
		if err != nil {
			return err
		}
		return WriteJSON(w, http.StatusOK, dailyWeatherList)
	}
	return fmt.Errorf("Method not allowed: %s", r.Method)
}

func (server *Server) handleGetDailyWeatherWithAuth(w http.ResponseWriter, r *http.Request) error {
	apiKey := GetApiKey()
	if r.Method == "GET" {
		if r.Header.Get("api-key") == apiKey {
			vars := mux.Vars(r)
			dailyWeatherList, err := server.repository.GetDailyWeatherWithAuth(vars["city"], vars["day"])
			if err != nil {
				return err
			}
			return WriteJSON(w, http.StatusOK, dailyWeatherList)
		} else {
			return fmt.Errorf("API Key is not valid.")
		}
	}
	return fmt.Errorf("Method not allowed: %s", r.Method)
}
