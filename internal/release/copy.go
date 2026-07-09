package release

import (
	"io"
	"os"
	"path/filepath"
)

func CopyFile(source, destination string) error {

	if err := os.MkdirAll(filepath.Dir(destination), os.ModePerm); err != nil {
		return err
	}

	src, err := os.Open(source)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)

	return err
}

func CopyDirectory(source, destination string) error {

	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		rel, err := filepath.Rel(source, path)
		if err != nil {
			return err
		}

		target := filepath.Join(destination, rel)

		if info.IsDir() {
			return os.MkdirAll(target, os.ModePerm)
		}

		return CopyFile(path, target)
	})
}
