package services

import (
	"blog_backend/dto"
	"blog_backend/properties"
	"fmt"
	"net/smtp"

	"go.mongodb.org/mongo-driver/bson"
)

// smtpServer data to smtp server
type smtpServer struct {
	host string
	port string
}

// Address URI to smtp server
func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

func TriggerEmail(user dto.UserDetails) int {
	if user.UserName == "" {
		return 3001
	}

	searchRequest := searchRequestBuilderForUserName(user.UserName)
	var userDetails dto.UserDetails
	bsonBytes, _ := bson.Marshal(executeSearch(properties.BLOG_BACKEND_DATABASE, properties.USER_DETAILS_COLLECTION, searchRequest))
	bson.Unmarshal(bsonBytes, &userDetails)

	if user.UserName != userDetails.UserName {
		return 3001
	}

	// Sender data.
	from := properties.SENDER_EMAIL
	password := properties.SENDER_PASSWORD
	// Receiver email address.
	to := []string{
		"gurjinders92571@gmail.com",
	}
	// smtp server configuration.
	smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpServer.host)

	var emailRequest dto.EmailRequestDTO
	emailRequest.To = to
	emailRequest.Subject = "[Important] Reset Propogate Password"

	status := TemplateService(&emailRequest, "ResetPassword.html", userDetails)

	if status == 3010 {
		return 30010
	}

	// Message.
	message := []byte("To: " + userDetails.Email + "m\r\n" +
		"Subject:" + emailRequest.Subject + "\r\n" +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" +
		"\r\n" +
		"" + emailRequest.Body + "\r\n")

	// Sending email.
	err := smtp.SendMail(smtpServer.Address(), auth, from, emailRequest.To, message)
	if err != nil {
		fmt.Println(err)
		return 3008
	}
	return 3009
}
