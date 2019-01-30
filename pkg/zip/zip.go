package zip

import (
	"archive/zip"
	"bytes"
	"io"
	"log"
	"os"
	"path/filepath"
)

// Extract writes all files at the destination path.
func Extract(data []byte, dest string) (uint64, error) {

	zr, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return 0, err
	}

	err = os.MkdirAll(dest, 0755)
	if err != nil {
		log.Printf("Cannot create extension folder: %s\n", dest)
	}

	writeFile := func(zf *zip.File, dest string) {

		// If this is a Directory, create it and move on.
		if zf.FileInfo().IsDir() {
			os.MkdirAll(dest, 0755)
			return
		}

		fr, err := zf.Open()
		if err != nil {
			log.Printf("Unable to read file: %s\n", zf.Name)
			return
		}
		defer fr.Close()

		path := filepath.FromSlash(filepath.Join(dest, zf.Name))
		dt := filepath.Dir(path)

		// Make the directory required by this File.
		os.MkdirAll(dt, 0755)

		// Create File.
		f, err := os.Create(path)
		if err != nil {
			log.Printf("Unable to create file: %s\n", path)
			return
		}

		defer f.Close()

		// Copy contents to the File.
		_, err = io.Copy(f, fr)
		if err != nil {
			log.Printf("Problem writing contents to file <%s>: %s\n", path, err)
			f.Close()
			return
		}

		err = f.Close()
		if err != nil {
			log.Printf("Problem closing file <%s>: %s\n", path, err)
			return
		}

		return

	}

	var size uint64

	// Create each File in the Archive.
	for _, zf := range zr.File {
		writeFile(zf, dest)
		size += zf.UncompressedSize64
	}

	return size, nil

}
