package logedUser

type LoggedUserRepository interface {
	GetUserAllInfo() LoggedUser

	SetUserEmail(email string)

	GetUserName() string
	GetUserEmail() string
}

type LoggedUser struct {
	Email string
	Name  string
}

//func ()  {
//
//}
