package conversion_utils

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/prashant42b/trademarks-elastic-search-engine/config"
)

func UnzipXML(zipFileName string, targetFolder string) {
	// Path to zip file
	zipFilePath := config.ZIP_PATH + zipFileName

	// XML file name
	xmlFileName := config.XML_NAME

	err := unzip(zipFilePath, xmlFileName, targetFolder)

	if err != nil {
		fmt.Println("Error unzipping file:", err)
		return
	}
}

func unzip(zipFilePath, fileName, targetFolder string) error {
	r, err := zip.OpenReader(zipFilePath)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, file := range r.File {
		if file.Name == fileName {
			rc, err := file.Open()
			if err != nil {
				return err
			}
			defer rc.Close()

			outFilePath := filepath.Join(targetFolder, fileName)
			outFile, err := os.Create(outFilePath)
			if err != nil {
				return err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, rc)
			if err != nil {
				return err
			}

			return nil
		}
	}

	return fmt.Errorf("file %s not found in the zip archive", fileName)
}
