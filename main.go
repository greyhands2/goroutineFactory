package main

import (
	"github.com/whatever/your/app/name/is/goroutines"
	"github.com/whatever/your/app/name/is/jobs"
	"github.com/whatever/your/app/name/is/model"
)

func main() {
	// define user
	user := &model.User{
		Email: "tonystarq@marvel.io",
		Image: "https:www.blablabla-image/images/1234/tony.png",
	}

	// define job data for deleting image
	imageData := jobs.DeleteImageJobData{
		ImageUrl: user.Image,
	}
	// call go routine factory to process and upload the image
	go goroutines.WorkGoroutine[jobs.DeleteImageJobData](jobs.DeleteImage, imageData)

	emailData := jobs.EmailJobData{
		User: user, Subject: "Maharba Mizterrr Wick", Body: "You know how it goes 😉😉😉 say the passcode and your weapons shall be yours🙇🏻🙇🏻🙇🏻 : ",
	}
	go goroutines.WorkGoroutine[jobs.EmailJobData](jobs.SendEmail, emailData)
}
