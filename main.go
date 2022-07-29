package main

import (
	"context"
	_ "embed"
	"github.com/pterm/pterm"
	"os"
	"path/filepath"
	"strings"
	"time"
	"github.com/pterm/pterm/putils"
)

//go:embed serviceAccountKey.json
var credentialsFile []byte

//go:embed .env
var envData []byte

//go:embed accounts.toml
var accountData []byte

var ctx = context.Background()

func main() {

	pterm.Println()
	pterm.DefaultBigText.WithLetters(putils.LettersFromStringWithStyle("MC-Sync-Upload", pterm.NewStyle(pterm.FgCyan))).Render()

	pterm.Println()

	pterm.Info.Println("Asegurate de el mundo esté ubicado en la carpeta correspondiente")

	pterm.Println()

	pterm.DefaultSection.Println("📚 Listando carpetas en el directorio")

	pterm.Println()

	var folders []string

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			if path != "." && !strings.Contains(path, "\\") {
				folders = append(folders,  path)
			}
		}
		return nil

	})

	if len(folders) == 0 {
		pterm.Error.Println("🚫 No se encontraron carpetas")
		return
	}

	var folderName string

	folderName, _ = pterm.DefaultInteractiveSelect.WithDefaultText("-> Seleccione el mundo").
	WithOptions(folders).
	Show()

	pterm.Println()

	pterm.Info.Println("Aplicando Configuraciones 🔨")

	err = setPlayersConfig(folderName)

	if err != nil {
		pterm.Fatal.Println(err)
	}

	var currentTime = time.Now().Format("2006-01-02")

	var zipFileName string = folderName + "-" + currentTime + ".zip"

	zipSource(folderName, "./"+zipFileName)

	fileData, err := os.ReadFile(zipFileName)

	if err != nil {
		pterm.Fatal.Println(err)
	}

	var option string

	var options = []string{
		"Subir a la nube ☁",
		"Mandar por url 🔗",
	};

	pterm.Println()
	option, _ = pterm.DefaultInteractiveSelect.WithDefaultText("-> Seleccione un método de descarga").
	WithOptions(options).
	Show()

	if option == options[0] {

		err = uploadWorld(zipFileName, fileData)

		if err != nil {
			pterm.Fatal.Println(err)
		}

		pterm.Println()

		pterm.Info.Println("📚 Subiendo archivo al servidor")

		pterm.Println()

		pterm.Success.Println("✨ Archivo subido correctamente ✨")

	} else if option == options[1] {

		err = serveWorld(zipFileName)

		if err != nil {
			pterm.Fatal.Println(err)
		}

	}

	os.Remove(zipFileName)

	time.Sleep(time.Second * 5)

	pterm.Info.Println("🏁 Saliendo del programa 🏁")

	os.Exit(0)

}
