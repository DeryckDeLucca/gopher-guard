package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DeryckDeLucca/gopher-guard.git/internal/monitor"
)

// Config define a estrutura do nosso arquivo JSON
type Config struct {
	Targets       []string `json:"targets"`
	CheckInterval int      `json:"check_interval"`
}

func loadConfig() Config {
	file, _ := os.ReadFile("config.json")
	var config Config
	json.Unmarshal(file, &config)
	return config
}

func main() {
	os.Mkdir("reports", 0755)
	config := loadConfig()

	// Canal para capturar o comando de fechar (Ctrl+C)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Ticker para rodar o monitoramento periodicamente
	ticker := time.NewTicker(time.Duration(config.CheckInterval) * time.Second)
	defer ticker.Stop()

	fmt.Println("🛡️ GopherGuard Pro Ativado")
	fmt.Printf("Monitorando %d alvos a cada %d segundos...\n", len(config.Targets), config.CheckInterval)

	runAudit(config.Targets) // Roda a primeira vez imediatamente

	for {
		select {
		case <-ticker.C:
			runAudit(config.Targets)
		case <-stop:
			fmt.Println("\nSinal de interrupção recebido. Encerrando GopherGuard com segurança...")
			return
		}
	}
}

func runAudit(targets []string) {
	resultsChan := make(chan monitor.Result, len(targets))
	var finalResults []monitor.Result

	for _, url := range targets {
		go monitor.CheckStatus(url, resultsChan)
	}

	for i := 0; i < len(targets); i++ {
		res := <-resultsChan
		finalResults = append(finalResults, res)
		if !res.Status {
			fmt.Printf("⚠️ [%s] ALERTA: %s está FORA DO AR!\n", time.Now().Format("15:04:05"), res.URL)
		}
	}

	monitor.SaveReport(finalResults)
	fmt.Printf("✅ Auditoria realizada em %s. Relatório atualizado.\n", time.Now().Format("15:04:05"))
}
