package main

import (
	"fmt"
	"os"

	"github.com/ttmday/go-logger-colorized/src/logger"
	"github.com/ttmday/gomailerterminal/src/config"
	"github.com/ttmday/gomailerterminal/src/repositories/mailer"
)

func main() {
	config.SetFiglet("Gomailer")

	var auth *mailer.MailerAuth

	switch os.Args[1] {
	case "-u":
		if len(os.Args) >= 5 {
			username := os.Args[2]
			password := os.Args[4]
			auth = mailer.LoadCredentialsFromFlags(username, password)
		} else {
			fmt.Println("-u email -p password")
			return
		}
	case "-f":
		if len(os.Args) >= 3 {
			filename := os.Args[2]
			a, err := mailer.LoadCredentialsFromFile(filename)
			if err != nil {
				logger.Error().Panicf("Error Cargando credenciales %v", err)
				return
			}

			auth = a
		}
	default:
		logger.Info().Println("Valores subministrados insuficientes")
		return
	}

	// m := mailer.New(&mailer.Mail{

	// }, auth)
	logger.Success().Printf("Credenciales cargadas %v", auth)
}
