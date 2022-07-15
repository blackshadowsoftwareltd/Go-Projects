package read_write

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
)

func WriteFile(file multipart.File) {
	tempFile, err := ioutil.TempFile("temp_file", "upload-*.png")
	if err != nil {
		fmt.Println("Error creating temp file")
		fmt.Println(err)
		return
	}
	defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file")
		fmt.Println(err)
		return
	}
	tempFile.Write(fileBytes)

	fmt.Println("Successfully wrote file to:", tempFile.Name())
}
