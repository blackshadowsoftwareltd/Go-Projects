package read_write

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"strings"
)

func WriteFile(file multipart.File, fileName string) (message string, success bool) {

	///? get file extension
	namesList := strings.Split(fileName, ".")
	extension := "png"
	if len(namesList) < 2 {
		fmt.Println("too short file name")
	} else {
		extension = namesList[len(namesList)-1]
	}
	fmt.Println("File extension : ", extension)

	///? read file
	tempFile, err := ioutil.TempFile("temp_file", namesList[0]+"*."+extension) //! everytime a new file will be stored inside temp_file directory.
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
	//!	fmt.Println("file bytes : ", fileBytes)
	///? It will write a file from the bytes under the temp_file directory.
	tempFile.Write(fileBytes)

	fmt.Println("Successfully wrote file to:", tempFile.Name())
	return "Successfully wrote file to :" + tempFile.Name(), true
}
