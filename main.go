package main

import (
	"os"

	"github.com/ttmday/go-logger-colorized/src/logger"
	"github.com/ttmday/gomailerterminal/src/config"
	"github.com/ttmday/gomailerterminal/src/repositories/mailer"
)

func main() {
	config.SetFiglet("Gomailer")

	var auth *mailer.MailerAuth

	switch os.Args[1] {
	case "--help":
		mailer.Usage()
		return
	case "-u":
		if len(os.Args) >= 5 {
			username := os.Args[2]
			password := os.Args[4]
			auth = mailer.LoadCredentialsFromFlags(username, password)
		} else {
			mailer.Usage()
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
		mailer.Usage()
		return
	}

	mail, err := mailer.Init()

	if err != nil {
		logger.Error().Panicf("Error en la obtencion de datos: %v", err)
		return
	}

	m := mailer.New(mail, auth)

	ms, err := m.CreateMail()

	if err != nil {
		logger.Error().Panicf("Error en la creación del correo: %v", err)
		return
	}

	_, err = ms.SendMail()

	if err != nil {
		logger.Error().Panicf("Error al enviar correo: %v", err)
		return
	}

	logger.Success().Println("Correo enviado.")

}
