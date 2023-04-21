package main

import (
	"os"

	"github.com/ttmday/go-logger-colorized/src/logger"
	"github.com/ttmday/gomailerterminal/src/config"
	"github.com/ttmday/gomailerterminal/src/global/helpers/args"
	"github.com/ttmday/gomailerterminal/src/global/helpers/utl"
	"github.com/ttmday/gomailerterminal/src/repositories/mailer"
)

func main() {
	config.SetFiglet("Gomailer")

	var auth *mailer.MailerAuth

	a := args.New(os.Args[1:])

	if a == nil {
		mailer.Usage()
		return
	}

	if a.IndexOf("--help") != -1 {
		mailer.Usage()
		return
	}

	html := ""
	var err error
	htmlIdx := a.IndexOf("--html")
	if htmlIdx != -1 {
		utl.Mapping(htmlIdx+1, os.Args[1:], func(s []string) {
			html, err = utl.LoadFile(s[0])
			if err != nil {
				logger.Error().Fatalf("Error Leyendo Contenido del Html %v", err)
				return
			}
		})

	}

	switch os.Args[1] {
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
				logger.Error().Fatalf("Error Cargando credenciales %v", err)
				return
			}

			auth = a
		} else {
			mailer.Usage()
			return
		}
	default:
		mailer.Usage()
		return
	}

	mail, err := mailer.Init(html)

	if err != nil {
		logger.Error().Fatalf("Error en la obtencion de datos: %v", err)
		return
	}

	m := mailer.New(mail, auth)

	ms, err := m.CreateMail()

	if err != nil {
		logger.Error().Fatalf("Error en la creación del correo: %v", err)
		return
	}

	_, err = ms.SendMail()

	if err != nil {
		logger.Error().Fatalf("Error al enviar correo: %v", err)
		return
	}

	logger.Success().Println("Correo enviado.")

}
