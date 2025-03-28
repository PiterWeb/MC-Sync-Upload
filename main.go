package main

import (
	_ "embed"

	"github.com/piterweb/mc-sync/src"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {

	pterm.Println()
	err := pterm.DefaultBigText.WithLetters(putils.LettersFromStringWithStyle("MC-Sync", pterm.NewStyle(pterm.FgCyan))).Render()

	if err != nil {
		pterm.Fatal.Println(err)
	}

	MODE_DOWNLOAD := "Descargar mundo"
	MODE_UPLOAD := "Subir mundo"

	MODES := []string{
		MODE_DOWNLOAD,
		MODE_UPLOAD,
	}

	mode, err := pterm.DefaultInteractiveSelect.WithDefaultText("-> Seleccione la ación").
		WithOptions(MODES).
		Show()

	if err != nil {
		pterm.Fatal.Println(err)
	}

	if mode == MODE_UPLOAD {
		src.UploadMain()
	} else if mode == MODE_DOWNLOAD {

	}

}
