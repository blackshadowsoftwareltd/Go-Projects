package uploads

import (
	message "FileUploading/messages"
	readWrite "FileUploading/read_write"
	"fmt"

	"net/http"
)

func FileUpload(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("image") //! image is a field. It will be use in the PostMan's Form's key.

	if err != nil {
		fmt.Println("Error retrieving the file")
		fmt.Println(err)
		message.ErrorMessage(w, r, "Error retrieving the file. Error : "+err.Error())
		return
	}
	defer file.Close()
	fmt.Println("File:", handler.Filename)
	fmt.Println("Size:", handler.Size)
	fmt.Println("MIME Header:", handler.Header)

	///? write file
	msg, success := readWrite.WriteFile(file, handler.Filename)
	if !success {
		message.ErrorMessage(w, r, msg)
	} else {
		message.SendMessage(w, r, msg)
	}
}
