package bin

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"os/exec"

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

	err := exec.Command("./"+programFolder+"/AnvilPacker", "pack", "-i", worldPath, "-o", outputName, "--preset", "smaller").Run()

	if err != nil {
		return nil, err
	}

	fileCompressed, err := os.ReadFile(outputName)

	if err != nil {
		return nil, err
	}

	return fileCompressed, nil
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
