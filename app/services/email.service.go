// using SendGrid's Go Library
// https://github.com/sendgrid/sendgrid-go
package services

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendWelcomeEmail(email string) error {
	from := mail.NewEmail("Artom Team", "hello@artom.io")
	subject := "Welcome to artom.io!"
	to := mail.NewEmail("", email)
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := `<div
	<p>Hi there !</p
		<p>Thanks for signing up and welcome to the artom.io Community!</p
		<p>At artom.io, we are building a community-drivern platform where you will be able to buy, sell and collect shares of iconic artworks. Soon, you will be able to invest in iconic blue-chip artworks curated and verified by our team of expert curators.</p
		<p>We also aim to change the vision of art. We see art as a good for all and we want to democratize it and make it popular. All the art that the community acquire will be lent and displayed in a museum of our network... and you get to decide which one!</p
		<p>We will be launching soon and we look forward to having you starting your collection</p
		<p>In the meantime, you can follow us on our social media or reach out <a href='https://www.artom.io'>here</a></p
		<p>The artom.io team,</p
	</div>`
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
		return err
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
		return nil
	}
}
