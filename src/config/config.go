package config

import (
	"fmt"
	"io"
	"os"

	"github.com/mbndr/figlet4go"
	"github.com/ttmday/go-logger-colorized/src/logger"
)

func SetFiglet(s string) {
	ascii := figlet4go.NewAsciiRender()

	options := figlet4go.NewRenderOptions()
	options.FontName = "JetBrains Mono"
	options.FontColor = []figlet4go.Color{
		// Colors can be given by default ansi color codes...
		figlet4go.ColorGreen,
	}

	renderStr, err := ascii.RenderOpts(s, options)

	if err == nil {
		fmt.Print(renderStr)
	}
}

// func

func LoadEnvFromFile(filename string) error {
	// LEEMOS EL ARCHIVO SUBMINISTRADO COMO ARGUMENTO //
	file, err := os.OpenFile(filename, os.O_RDWR, 0o600)
	if err != nil {
		logger.Error().Println("Error al leer archivo | LoadEnvFromFile ", err)
		return err
	}

	// CERRAMOS AL RETORNAR //
	defer file.Close()

	// OBTENEMOS LA INFORMACION DEL ARCHIVO LEIDO //
	fileInfo, err := file.Stat()
	if err != nil {
		logger.Error().Println("Error al leer información del archivo | LoadEnvFromFile ", err)
		return err
	}

	// VALIDAMOS SU TAMAÑO Y NOMBRE //

	if fileInfo.Size() < 0 {
		logger.Error().Println("El Archivo cargado no posee información | LoadEnvFromFile ", err)
		return err
	}

	if fileInfo.Name() != ".env" {
		logger.Error().Println("El Nombre del archivo debe ser .env | LoadEnvFromFile")
		return err
	}

	// FIN DE LA VALIDACION //

	// CARGAMOS EL ARCHIVO DE ENVIROMENTS //
	fileEnv, err := os.OpenFile("./.env", os.O_RDWR|os.O_CREATE, 0o600)
	if err != nil {
		logger.Error().Println("Error al leer archivo | LoadEnvFromFile ", err)
		return err
	}

	// CERRAMOS AL RETORNAR //
	defer fileEnv.Close()

	// COMPIAMOS LOS DATOS DEL ARCHIVO SUBMINISTRADO AL ARCHIVO DE ENVIROMENTS //
	_, err = io.Copy(fileEnv, file)

	if err != nil {
		logger.Error().Println("Error al copiar archivo en el .env | LoadEnvFromFile ", err)
		return err
	}

	return nil
}
