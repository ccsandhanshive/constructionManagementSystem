package services

type Main interface {
	login(userid int, password string) string                         //login user
	changePass(username int, password string, newpassword string) int //change in password
}
