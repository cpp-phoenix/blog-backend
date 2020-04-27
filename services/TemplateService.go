package services

import (
	"blog_backend/dto"
	"blog_backend/properties"
	"bytes"
	"fmt"
	"html/template"
	"os"
)

func TemplateService(request *dto.EmailRequestDTO, templateName string, user dto.UserDetails) int {
	var resetPassword dto.ResetPassword
	resetPassword.UserName = user.UserName
	resetPassword.ConfirmEmailUrl = properties.FRONT_END_BASE_URL + "/resetPassword?" + GenerateUrl(user)
	resetPassword.DeactivateAccountUrl = properties.FRONT_END_BASE_URL + "/deactivateAccount?" + GenerateUrl(user)
	basepath, _ := os.Getwd()
	t, err := template.ParseFiles(basepath + "/templates/" + templateName)
	if err != nil {
		fmt.Println(err)
		return 3010
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, resetPassword); err != nil {
		fmt.Println(err)
		return 3010
	}
	request.Body = buf.String()
	return 3011
}
