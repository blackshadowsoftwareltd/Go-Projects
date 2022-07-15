package read_write

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
)

func WriteFile(file multipart.File) (message string, success bool) {
	tempFile, err := ioutil.TempFile("temp_file", "upload-*.png")
	if err != nil {
		fmt.Println("Error creating temp file")
		fmt.Println(err)
		message = "Error creating temp file. Error : " + err.Error()
		return message, false
	}
	defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file")
		fmt.Println(err)
		message = "Error reading file. Error : " + err.Error()
		return message, false
	}
	tempFile.Write(fileBytes)

	fmt.Println("Successfully wrote file to:", tempFile.Name())
	return "Successfully wrote file to:" + tempFile.Name(), true
}
