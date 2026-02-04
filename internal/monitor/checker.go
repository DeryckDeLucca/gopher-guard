package monitor

import (
	"net/http"
	"time"
)

type Result struct {
	URL     string
	Status  bool
	Latency time.Duration
	Quality string // Estável, Lento ou Crítico
}

func CheckStatus(url string, c chan Result) {
	start := time.Now()
	client := http.Client{Timeout: 5 * time.Second}

	resp, err := client.Get(url)
	latency := time.Since(start)

	if err != nil || resp.StatusCode >= 400 {
		c <- Result{URL: url, Status: false, Latency: latency, Quality: "CRÍTICO"}
		return
	}

	// Lógica de classificação de qualidade
	quality := "ESTÁVEL"
	if latency > 500*time.Millisecond {
		quality = "LENTO"
	}
	if latency > 1500*time.Millisecond {
		quality = "INSTÁVEL"
	}

	c <- Result{URL: url, Status: true, Latency: latency, Quality: quality}
}
