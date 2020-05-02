package dto

type UserDetails struct {
	Name                 string
	UserName             string
	Email                string
	Password             string
	Otp                  int
	Avatar               int
	CreatedPosts         []int
	ResetRequestTimstamp int64
	CreatedTimestamp     int64
}
