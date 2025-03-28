package bin

import (
	"embed"
	"fmt"
	"io/fs"
	"os"

	// "os/exec"
	"path/filepath"
)

const (
	OpPack = iota
	OpUnpack
)

const programFolder = "AnvilPacker"

func UseAnvilPacker(worldPath string, outputName string, op int) ([]byte, error) {

	defer os.RemoveAll(programFolder)

	err := copyDirFromEmbed(anvilPackerFolder, anvilPackerName, programFolder)

	if err != nil {
		return nil, err
	}

	if op == OpPack {
		return UseAnvilPackerPack(worldPath, outputName)
	} else if op == OpUnpack {
		return UseAnvilPackerUnpack(worldPath, outputName)
	} else {
		return nil, fmt.Errorf("Invalid operation")
	}

}

func UseAnvilPackerPack(worldPath string, outputName string) ([]byte, error) {

	// exec.Command("./" + programFolder + "/AnvilPacker", "-i", worldPath, "-o", outputName).Run()

	return []byte{}, fmt.Errorf("AnvilPackerPack not implemented")
}

func UseAnvilPackerUnpack(worldPath string, outputName string) ([]byte, error) {
	return []byte{}, fmt.Errorf("AnvilPackerUnpack not implemented")
}

func copyDirFromEmbed(fsys embed.FS, root string, destDir string) error {
	return fs.WalkDir(fsys, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(destDir, relPath)

		if d.IsDir() {
			return os.MkdirAll(destPath, 0o700)
		}

		data, err := fs.ReadFile(fsys, path)
		if err != nil {
			return err
		}

		return os.WriteFile(destPath, data, 0o644)
	})
}
