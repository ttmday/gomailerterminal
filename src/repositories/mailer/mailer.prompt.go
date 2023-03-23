package mailer

import (
	"errors"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/ttmday/gomailerterminal/src/global/helpers"
)

func Init() (*Mail, error) {

	to, err := promptTo()
	if err != nil {
		return nil, err
	}

	dstName, err := promptDstName()
	if err != nil {
		return nil, err
	}

	subject, err := promptSubject()
	if err != nil {
		return nil, err
	}

	message, err := promptMessage()
	if err != nil {
		return nil, err
	}

	return &Mail{
		To:      strings.TrimSpace(to),
		DstName: strings.TrimSpace(dstName),
		Subject: strings.TrimSpace(subject),
		Message: strings.TrimSpace(message),
	}, nil
}

func promptTo() (string, error) {
	prompt := promptui.Prompt{
		Label: "Para",
		Validate: func(s string) error {
			if helpers.IsAEmail(s) == false {
				return errors.New("El valor debe ser un correo eléctronico válido.")
			}

			return nil
		},
	}

	return prompt.Run()
}

func promptDstName() (string, error) {
	prompt := promptui.Prompt{
		Label: "Destinatario",
		Validate: func(s string) error {
			if s == "" {
				return errors.New("El valor no debe ser vacio.")
			}
			return nil
		},
	}

	return prompt.Run()
}

func promptSubject() (string, error) {
	prompt := promptui.Prompt{
		Label: "Asunto",
		Validate: func(s string) error {
			if s == "" {
				return errors.New("El valor no debe ser vacio.")
			}
			return nil
		},
	}

	return prompt.Run()
}

func promptMessage() (string, error) {
	prompt := promptui.Prompt{
		Label: "Mensaje",
		Validate: func(s string) error {
			if s == "" {
				return errors.New("El valor no debe ser vacio.")
			}
			return nil
		},
	}

	return prompt.Run()
}
