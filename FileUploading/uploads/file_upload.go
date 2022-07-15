package uploads

import (
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
		return
	}
	defer file.Close()
	fmt.Println("File:", handler.Filename)
	fmt.Println("Size:", handler.Size)
	fmt.Println("MIME Header:", handler.Header)

	readWrite.WriteFile(file)

	// fmt.Print("File Upload Endpoint Hit\n\n")
	// // Get the file from the request
	// file, header, err := r.FormFile("file")
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }
	// defer file.Close()

	// // Create a new file
	// out, err := os.Create("uploads/" + header.Filename)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// defer out.Close()

	// // Write the file to the new file
	// _, err = io.Copy(out, file)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// w.WriteHeader(http.StatusOK)
}
