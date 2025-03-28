package src

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/piterweb/mc-sync/src/bin"
	"github.com/pterm/pterm"
)

func UploadMain() {

	pterm.Println()

	pterm.Info.Println("Asegurate de el mundo esté ubicado en la carpeta correspondiente")

	pterm.Println()

	pterm.DefaultSection.Println("📚 Listando carpetas en el directorio")

	pterm.Println()

	folders := []string{}

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

	if err != nil {
		pterm.Fatal.Println(err)
	}

	if len(folders) == 0 {
		pterm.Error.Println("🚫 No se encontraron carpetas")
		return
	}

	folderName, err := pterm.DefaultInteractiveSelect.WithDefaultText("-> Seleccione el mundo").
		WithOptions(folders).
		Show()

	if err != nil {
		pterm.Fatal.Println(err)
	}

	pterm.Println()

	pterm.Info.Println("Aplicando Configuraciones 🔨")

	rawConfig, err := os.ReadFile("./" + filepath.Join(folderName, "config.toml"))

	if err != nil {
		pterm.Fatal.Println(err)
	}

	err = SetPlayersConfig(folderName, rawConfig)

	if err != nil {
		pterm.Fatal.Println(err)
	}

	var currentTime = time.Now().Format("2006-01-02")

	compressedFileName := folderName + "-" + currentTime + ".apw"

	fileData, err := bin.UseAnvilPacker(folderName, compressedFileName, bin.OpPack)

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

		err = UploadWorld(rawConfig, compressedFileName, fileData)

		if err != nil {
			pterm.Fatal.Println(err)
		}

		pterm.Println()

		pterm.Info.Println("📚 Subiendo archivo al servidor")

		pterm.Println()

		pterm.Success.Println("✨ Archivo subido correctamente ✨")

	} else if option == options[1] {

		err = ServeWorld(compressedFileName)

		if err != nil {
			pterm.Fatal.Println(err)
		}

	}

	os.Remove(compressedFileName)

	time.Sleep(time.Second * 5)

	pterm.Info.Println("🏁 Saliendo del programa 🏁")

	os.Exit(0)

}
