package jobs

import (
	"github.com/whatever/your/app/name/is/image"
)

type DeleteImageJobData struct {
	ImageUrl string
}

func DeleteImage(data DeleteImageJobData) (result string, message string, error error) {
	// send to a goroutine and return a response to the user

	//return "Success", "Success", nil
	_, err := image.DeleteFromCloudinary(data.ImageUrl)

	if err != nil {
		return "Failed", "Error Deleting Image On Cloudinary", err
	}

	return "Success", "Image Successfully Deleted", nil

}
