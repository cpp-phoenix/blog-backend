package dto

type UserDetails struct {
	UserName             string
	Email                string
	Password             string
	Otp                  int
	ResetRequestTimstamp int64
}
