package bin

import (
	"embed"
)

const anvilPackerName = "AnvilPacker-0.9.9-beta-win-x64"

//go:embed AnvilPacker-0.9.9-beta-win-x64
var anvilPackerFolder embed.FS
