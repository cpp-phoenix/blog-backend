package dto

type EmailRequestDTO struct {
	To      []string
	Subject string
	Body    string
}
