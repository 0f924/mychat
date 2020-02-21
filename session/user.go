package session

type User struct {
	UserId   string
	UserName string
}

func (this User) String() string {
	return this.UserName + "(" + this.UserId + ")"
}
