package monitor

import (
	"fmt"
	"os"
	"time"
)

func SaveReport(results []Result) error {
	fileName := fmt.Sprintf("reports/audit_%s.md", time.Now().Format("2006-01-02_15-04"))

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	content := "# 🛡️ Relatório de Auditoria GopherGuard\n\n"
	content += fmt.Sprintf("Data da verificação: %s\n\n", time.Now().Format(time.RFC1123))
	content += "| Dispositivo/URL | Status | Latência | Qualidade |\n"
	content += "| :--- | :--- | :--- | :--- |\n"

	for _, res := range results {
		status := "✅ OK"
		if !res.Status {
			status = "❌ FALHA"
		}
		content += fmt.Sprintf("| %s | %s | %v | %s |\n", res.URL, status, res.Latency.Round(time.Millisecond), res.Quality)
	}

	_, err = file.WriteString(content)
	return err
}
