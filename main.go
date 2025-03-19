package main

import (
	"context"
	_ "embed"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/piterweb/mc-sync/src"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

var ctx = context.Background()

func main() {

	pterm.Println()
	pterm.DefaultBigText.WithLetters(putils.LettersFromStringWithStyle("MC-Sync", pterm.NewStyle(pterm.FgCyan))).Render()

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
				folders = append(folders, path)
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

	rawConfig, err := os.ReadFile("./config.toml")

	if err != nil {
		pterm.Fatal.Println(err)
	}

	err = src.SetPlayersConfig(folderName, rawConfig)

	if err != nil {
		pterm.Fatal.Println(err)
	}

	var currentTime = time.Now().Format("2006-01-02")

	var zipFileName string = folderName + "-" + currentTime + ".zip"

	src.ZipSource(folderName, "./"+zipFileName)

	fileData, err := os.ReadFile(zipFileName)

	if err != nil {
		pterm.Fatal.Println(err)
	}

	var option string

	var options = []string{
		"Subir a la nube ☁",
		"Mandar por url 🔗",
	}

	pterm.Println()
	option, _ = pterm.DefaultInteractiveSelect.WithDefaultText("-> Seleccione un método de descarga").
		WithOptions(options).
		Show()

	if option == options[0] {

		err = src.UploadWorld(rawConfig, zipFileName, fileData)

		if err != nil {
			pterm.Fatal.Println(err)
		}

		pterm.Println()

		pterm.Info.Println("📚 Subiendo archivo al servidor")

		pterm.Println()

		pterm.Success.Println("✨ Archivo subido correctamente ✨")

	} else if option == options[1] {

		err = src.ServeWorld(zipFileName)

		if err != nil {
			pterm.Fatal.Println(err)
		}

	}

	os.Remove(zipFileName)

	time.Sleep(time.Second * 5)

	pterm.Info.Println("🏁 Saliendo del programa 🏁")

	os.Exit(0)

}
