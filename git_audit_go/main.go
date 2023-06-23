package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	horarioInicio = 10
	horarioFim    = 20
	since         = "6 months"
	dateFormat    = "2006-01-02 15:04:05 -0700"
	outputFile    = "output.csv"
)

func main() {
	// Executar o comando git log para obter os commits
	sinceParam := fmt.Sprintf("--since='%s'", since)
	cmd := exec.Command("git", "log", sinceParam, "--pretty=format:%H|%ai|%an|%ae")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Erro ao executar o comando git log: %v", err)
	}

	// Criar o arquivo CSV e escrever o cabeçalho
	file, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Erro ao criar o arquivo CSV: %v", err)
	}
	defer file.Close()

	// Cabeçalho do arquivo CSV
	csvHeader := "Hash,Data-Hora,Autor Nome,Autor E-mail"
	fmt.Fprintln(file, csvHeader)

	// Processar os commits
	commits := strings.Split(string(output), "\n")
	for _, commit := range commits {
		fields := strings.Split(commit, "|")
		if len(fields) < 4 {
			continue
		}

		hash := fields[0]
		dataHora := fields[1]
		autorNome := fields[2]
		autorEmail := fields[3]

		// Converter a data e hora para o tipo time.Time
		t, err := time.Parse(dateFormat, dataHora)
		if err != nil {
			log.Printf("Warning: Ignoring Commit with Invalid Date format: %s\n", commit)
			continue
		}

		// Converter para o fuso horário de Brasília
		t = t.In(time.FixedZone("BRT", -3*60*60))

		// Obter a hora do commit
		hourCommit := t.Hour()

		// Verificar se o commit ocorreu fora do horário de trabalho
		if hourCommit < horarioInicio || hourCommit >= horarioFim {
			// Escrever o commit no arquivo CSV
			line := fmt.Sprintf("%s,%s,%s,%s", hash, dataHora, autorNome, autorEmail)
			fmt.Printf("[FOUND] %s\n", line)
			fmt.Fprintln(file, line)
		}
	}

	fmt.Printf("CSV generated successfully: %s\n", outputFile)
}
