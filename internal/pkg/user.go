package pkg

type User struct {
	ID    uint
	Login string
	Role  string
}

var instance *User

func GetUser() *User {
	if instance == nil {
		instance = &User{
			ID:    1,
			Login: "fixed_creator",
			Role:  "creator",
		}
	}
	return instance
}