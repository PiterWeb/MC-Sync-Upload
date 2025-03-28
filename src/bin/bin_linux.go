package bin

import "embed"

const anvilPackerName = "AnvilPacker-0.9.9-beta-linux-x64"

//go:embed AnvilPacker-0.9.9-beta-linux-x64
var anvilPackerFolder embed.FS
