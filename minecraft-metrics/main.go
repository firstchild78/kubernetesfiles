package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var currentPlayerNo = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "current_player_number",
		Help: "Current number of players",
	},
)

var maxPlayerNo = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "max_player_number",
		Help: "Maximum numbrer of players",
	},
)

func main() {

	prometheus.MustRegister(currentPlayerNo)
	prometheus.MustRegister(maxPlayerNo)

	go func() {
		for {
			Init("localhost", "25565")
			current, err := strconv.ParseFloat(Current_players, 64)
			if err != nil {
				current = 0
			}

			max, err := strconv.ParseFloat(Max_players, 64)
			if err != nil {
				max = 0
			}

			currentPlayerNo.Set(current)
			maxPlayerNo.Set(max)
			time.Sleep(5 * time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":1234", nil)

}
