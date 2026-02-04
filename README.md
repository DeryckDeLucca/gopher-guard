# 🛡️ GopherGuard Pro

O **GopherGuard** é um sistema de auditoria de resiliência de rede desenvolvido em **Go**. Ele utiliza o poder da concorrência nativa (Goroutines e Channels) para monitorar a integridade de múltiplos ativos de rede simultaneamente, gerando relatórios detalhados de latência e qualidade.

## 🚀 Por que este projeto é eficiente?
- **Concorrência Real:** Diferente de scripts sequenciais, o GopherGuard verifica todos os alvos ao mesmo tempo.
- **Análise de Qualidade:** Classifica a saúde da conexão em: *Estável, Lento, Instável ou Crítico*.
- **Baixo Consumo:** Desenvolvido para rodar como um serviço leve de fundo.
- **Persistência em Markdown:** Gera logs automáticos em `.md` para auditoria técnica.

## 🛠️ Tecnologias
- **Language:** Go (Golang) 1.2x
- **Concorrência:** Goroutines & Channels
- **Formato de Dados:** JSON para configurações

## 🏗️ Estrutura
- `internal/monitor`: Core da lógica de rede e geração de relatórios.
- `reports/`: Histórico de auditorias em formato Markdown.
- `config.json`: Gerenciamento dinâmico de alvos.

## ⚙️ Como Usar
1. Adicione seus IPs/URLs no arquivo `config.json`.
2. Execute o serviço:
   ```bash
   go run main.go
3. Para interromper, use Ctrl+C. O programa realizará um Graceful Shutdown.