package users

type User struct {
	id    int
	Name  string
	Email string
}

func NewUser(id int, name string, email string) *User {
	ret := new(User)

	ret.id = id
	ret.Name = name
	ret.Email = email

	return ret
}

func (user *User) GetId() int {
	return user.id
}
