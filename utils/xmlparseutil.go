package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	//"encoding/json"
	//"encoding/xml"
)

func UnzipXML() {

	// Path to zip file
	zipFilePath := "/Users/anandsure/Desktop/USPTO data/apc18840407-20221231-01.zip"

	// XMl file name
	xmlFileName := "apc18840407-20221231-01.xml"

	// Path to extract zip file
	//extractPath := "/Users/anandsure/Desktop/USPTO data/extracted"

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
