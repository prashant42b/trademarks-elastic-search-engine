package conversion_utils

import (
	"archive/zip"
	"fmt"
	"io"
	"os"

	"github.com/prashant42b/elastic-search-engine-task/config"
	//"encoding/json"
	//"encoding/xml"
)

func UnzipXML(zipFileName string) {

	// Path to zip file
	zipFilePath := config.ZIP_PATH + zipFileName

	// XMl file name
	xmlFileName := config.XML_NAME

	err := unzip(zipFilePath, xmlFileName)

	if err != nil {
		fmt.Println("Error unzipping file:", err)
		return
	}

}

func unzip(zipFilePath, fileName string) error {

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

			outFile, err := os.Create(fileName)
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
