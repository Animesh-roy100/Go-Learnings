package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const api_key = "67b69cbdb9750df5fa7f62c0d45c535b"

func fetchWeatherOfCity(city string, ch chan<- string, wg *sync.WaitGroup) interface{} {
	var data struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}

	defer wg.Done()

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, api_key)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error fetching weather for %s: %v", city, err)
		return data
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Printf("error decoding weather data for %s: %v", city, err)
		return data
	}

	ch <- fmt.Sprintf("This is the %s", city)

	return data
}

func main() {
	startNow := time.Now()

	cities := []string{"Mumbai", "Delhi", "London", "Kolkata"}

	ch := make(chan string)
	var wg sync.WaitGroup

	for _, city := range cities {
		wg.Add(1)
		go fetchWeatherOfCity(city, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println(result)
	}

	fmt.Println("This operation took", time.Since(startNow))
}
